build:
	go mod download
	go build ./cmd/tivit-sftp

run-server:
	go run ./cmd/tivit-sftp/main.go server

run-client:
	go run ./cmd/tivit-sftp/main.go client

.PHONY: build run
