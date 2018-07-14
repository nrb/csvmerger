.DEFAULT_GOAL := build

test:
	go test ./pkg/...

build:
	go build -o csvmerger main.go 

all: test build

clean:
	rm csvmerger
