.PHONY: build test lint fmt build-workflow

build:
	go build ./...

test:
	go test ./...

lint:
	@out=$$(gofmt -l .); \
	if [ -n "$$out" ]; then \
		echo "gofmt needs to be run on:"; \
		echo "$$out"; \
		exit 1; \
	fi
	go vet ./...

fmt:
	gofmt -w .

build-workflow:
	scripts/build-workflow.sh
