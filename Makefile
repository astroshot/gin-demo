.PHONY: build
build:
	go build -v -o bin/gintama_api cmd/web/main.go

.PHONY: clean
clean:
	-rm bin/*

.PHONY: format
format:
	gofmt -w pkg cmd

.PHONY: start
start:
	bin/gintama_api
