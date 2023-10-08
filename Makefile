BINARY_NAME=pacman
build:
	@go build -o bin/$(BINARY_NAME) -v
	
run: build
	@./bin/$(BINARY_NAME)