LINTER_VER ?= v1.56.2-alpine

lint:
	golangci-lint run

test:
	go test -v -cover ./...
