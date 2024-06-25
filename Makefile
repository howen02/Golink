build:
	@go build -o bin/main

run: build
	@./bin/main

clean:
	@rm -rf bin

test:
	@go test