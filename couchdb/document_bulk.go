package couchdb

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

const (
	maxInsertByReq  int = 100
	maxParallelCall int = 8
)

type CreateDocumentsRequest struct {
	Docs []interface{} `json:"docs"`
}

// CreateDocuments creates documents
func CreateDocuments(ctx context.Context, documents []interface{}, databaseName string) error {
	if len(documents) == 0 {
		return nil
	}

	sem := semaphore.NewWeighted(int64(maxParallelCall))
	var g errgroup.Group

	for i := 0; i < len(documents); i += maxInsertByReq {
		i := i

		err := sem.Acquire(ctx, 1)
		if err != nil {
			return err
		}

		endIndex := i + maxInsertByReq
		if endIndex > len(documents) {
			endIndex = len(documents)
		}

		g.Go(func() error {
			defer sem.Release(1)
			log.Printf("batching [%d:%d]", i, endIndex)
			batch := documents[i:endIndex]
			_, err := PerformRequest(
				ctx,
				http.MethodPost,
				fmt.Sprintf("%s/_bulk_docs", databaseName),
				CreateDocumentsRequest{
					Docs: batch,
				},
			)

			if err != nil {
				return err
			}

			log.Printf(`successfully created documents [%d:%d]`, i, endIndex)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
