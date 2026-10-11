.PHONY: build test clean

build:
	bash build.sh

test:
	go test ./... -timeout 120s

clean:
	rm -f shellings shellings.exe shellings-darwin-arm64 shellings-darwin-amd64
