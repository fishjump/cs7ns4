.PHONY: all build

all: build

build:
	go build -o build/cs7ns4 cmd/main.go
	cp -r resource build

run:
	build/cs7ns4