GO ?= go
HAS_GO = $(shell hash $(GO) > /dev/null 2>&1 && echo "GO" || echo "NOGO")
SQLITE3 ?= sqlite3
HAS_SQLITE3 = $(shell hash $(SQLITE3) > /dev/null 2>&1 && echo "SQLITE3" || echo "NOSQLITE3")

ifeq ($(OS), Windows_NT)
	GOFLAGS := -v -buildmode=exe
	EXECUTABLE ?= nasu.exe
else ifeq ($(OS), Windows)
	GOFLAGS := -v -buildmode=exe
	EXECUTABLE ?= nasu.exe
else
	GOFLAGS := -v
	EXECUTABLE ?= nasu
endif

.PHONY: all

all: check-env run

check-env:
	@echo "checking go..."
ifeq ($(HAS_GO), NOGO)
	@echo "GO is not installed"
	exit -1
else
	@echo "GO is installed"
endif
	@echo "checking sqlite3..."
ifeq ($(HAS_SQLITE3), NOSQLITE3)
	@echo "sqlite3 is not installed"
	exit -1
else
	@echo "sqlite3 is installed"
endif

build: check-env
	go build -o $(EXECUTABLE) github.com/littlebutt/nasu/src

test:
	go test -cover ./...

run: check-env
	go run github.com/littlebutt/nasu/src

