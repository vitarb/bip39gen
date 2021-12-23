all: fmt build

BINARY_NAME=bip39gen

init:
	@git submodule update --init bips

fmt:
	@go fmt

build:
	go build -o ${BINARY_NAME} main.go
 
run: init build
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}
 
clean:
	@go clean
	@rm ${BINARY_NAME}
