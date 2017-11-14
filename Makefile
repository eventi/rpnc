.PHONY: test

test: rpnc
	go test
	bats -t ./test/*.bats

rpnc: *.go */*.go
	go build
