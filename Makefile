.PHONY: build run clean

BINARY_NAME=tgbots

build:
	go build -o $(BINARY_NAME) ./cmd/tgbots

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

lint:
	golangci-lint run
