.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go


run: build
	./.bin/bot

format:
	${call colored, formatting is running...}
	go vet ./...
	go fmt ./...
