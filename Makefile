# This is how we want to name the binary output
BINARY=bora

all:
	go build -o ${BINARY}

test:
	go test

clean:
	rm -f bora

get-deps:
	go get ./...

