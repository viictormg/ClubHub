run:
	go run cmd/main.go

test:
	go test ./... -cover -coverprofile=coverage.out

gen-mocks:
	mockery --dir=./internal --all


watch-coverage:
	go tool cover -html=coverage.out