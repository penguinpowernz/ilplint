.PHONY: build
build:
	go build ./cmd/ilplint

build_all:
	GOARCH=arm go build -o build/ilplint.arm ./cmd/ilplint
	GOARCH=arm64 go build -o build/ilplint.arm64 ./cmd/ilplint
	GOARCH=386 go build -o build/ilplint.386 ./cmd/ilplint
	GOARCH=amd64 go build -o build/ilplint.amd64 ./cmd/ilplint

clean:
	rm -rf build