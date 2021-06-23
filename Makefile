.PHONY:
.SILENT:
.DEFAULT_GOAL := run

run:
	go run cmd/main.go

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage

test.coverage:
	go tool cover -func=cover.out

migration-up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go

lint:
	golangci-lint run