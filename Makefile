all: test run

build:
	go build github.com/littlebutt/nasu/src

test:
	go test -cover ./...

run:
	go run github.com/littlebutt/nasu/src

