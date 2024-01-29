build:
	@go build -o bin/cache-go
run: build
	./bin/cache-go