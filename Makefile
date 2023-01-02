
couchdb_instance:
	docker run -d -e COUCHDB_USER=admin -e COUCHDB_PASSWORD=password -p 5984:5984 --name myCouchdb couchdb:latest
run:
	go run main.go

build:
	go build ./...

test:
	go clean -testcache && \
	go test ./...
