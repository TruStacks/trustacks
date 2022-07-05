.PHONY: test build clean

test:
	@go test ./... -v -race

build:
	@CGO_ENABLED=0 go build -o tsctl ./cmd

clean:
	@rm tsctl