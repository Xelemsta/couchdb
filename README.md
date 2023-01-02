## Prerequisites

* docker
* go

## About application

### Launch a couchdb instance

  ```sh
  make couchdb_instance
  ```

### Build application

  ```sh
  make build
  ```

### Start application

  ```sh
  make run
  ```

### Test application

  ```sh
  make test
  ```

### Extensions/Improvements

* Retrieve HTTP directly from ENV variable or by parameter of main.go.

* use HTTPS client.

* Store password safely (vault).

* Create configuration test file in /tmp (or any other suitable directory for tests purpose).

* Add more application logs, including but not limited to: response body, response time.

* Build a dedicated couchdb instance to perform tests (app and test are using the same instance).