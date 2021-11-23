.PHONY: all build

all: build

build:
	go build -o build/cs7ns4 cmd/main.go

run:
	build/cs7ns4