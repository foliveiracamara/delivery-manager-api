run:
	go run cmd/main.go server start

doc:
	swag init -g cmd/main.go
