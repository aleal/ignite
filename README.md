# Ignite POC
ignite with go 1.18 poc

## Run 

### Sample code
```shell
go run sample/example.v1/main.go
```

### Test coverage
```shell
go test -coverprofile /tmp/cp.out ./... && go tool cover -html=/tmp/cp.out
```