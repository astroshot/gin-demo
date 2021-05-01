# Project Info
VERSION := "1.0.0"
HOMEDIR := $(shell pwd)
BUILD_TARGET := $(HOMEDIR)/bin/gen-api
BUILD_SOURCE := $(HOMEDIR)/cmd/api/main.go
BUILD_INFO := "-X 'main.Version=$(VERSION)' \
		-X 'main.BuildTime=$(shell date "+%Y-%m-%d %H:%M:%S")' \
		-X 'main.GoVersion=$(shell go version | grep -Eo "go[0-9]+.* " | sed "s/ //")' \
		-X 'main.GitCommit=$(shell git rev-parse --short HEAD || echo unsupported)'"

build:
	go build -ldflags $(BUILD_INFO) -v -o $(BUILD_TARGET) $(BUILD_SOURCE)

clean:
	-rm bin/*

format:
	gofmt -w pkg cmd

start:
	bin/gin_api

.PHONY: build clean format start

