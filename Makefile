.PHONY: clean build bench unittest

build:
	go build -i ./...

unittest:
	go test ./...

bench:
	go test -bench ./...

clean:
	go clean -cache

