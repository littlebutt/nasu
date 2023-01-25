NPM ?= npm
HAS_NPM = $(shell hash $(NPM) > /dev/null 2>&1 && echo "NPM" || echo "NONPM")

.PHONY: all

all:
	@if [ -e ../src/web ]; then
	@echo "dir exists"

check-env:
	@echo "checking node..."
ifeq ($(HAS_NPM), NONPM)
	@echo "node is not installed"
	exit -1
else
	@echo "node is installed"
endif

rebuild: clean check-env
	npm run build
	mv build/index.html build/index.htm #go default behaviour
	mv build ../src/web

build:
	test -d ../src/web && echo "do not need to build" || make rebuild

lint: check-env
	npx eslint src/** --fix

clean:
	rm -rf ../src/web

run: check-env
	npm run start

