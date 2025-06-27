run:
	go run cmd/main.go server start

doc:
	swag init -g cmd/main.go

test:
	go test ./internal/... -v

test-coverage:
	go test ./internal/... -v -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

test-race:
	go test ./internal/... -v -race
