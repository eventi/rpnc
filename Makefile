.PHONY: test

test:
	bats ./test/*.bats

build: rpnc.go
	go build
