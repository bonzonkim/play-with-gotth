APP_NAME ?= app
build:
	go build -o ./bin/$(APP_NAME) ./cmd/main.go
