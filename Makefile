.PHONY: build clean format start
build:
	go build -v -o bin/gin_api cmd/web/server.go

clean:
	-rm bin/*

format:
	gofmt -w pkg cmd

start:
	bin/gin_api
