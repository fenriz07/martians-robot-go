# How use this

1. docker build -t martians .
2. docker run martians

## Test
You just do this: go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
