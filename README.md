# How use that

1. docker build -t martians .
2. docker run martians

## Test
You just do that: go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
