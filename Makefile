.PHONE: all, test, coverage

all: test

test:
	go fmt ./...
	go clean -testcache
	go test -coverprofile test.coverage

coverage:
	go tool cover -html test.coverage
