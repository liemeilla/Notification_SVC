run_api:
	go build ./cmd/api && ./api

run_test:
	go test ./internal/logic/...