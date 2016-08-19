# go-scylla-api

go-scylla-api is a Go client library for accessing the [Scylla REST API](https://github.com/scylladb/scylla/tree/master/api/api-doc).

Please note that this implementation is work-in-progress and does not cover the whole Scylla REST API.

## Usage

```go
import "github.com/penberg/go-scylla-api/scylla"
```

Create a new Scylla client:

```go
host := "localhost"
port := "10000"
client := scylla.NewClient(host, port)
```

Execute an API command using the `client` object:

```go
client.Flush()
```
