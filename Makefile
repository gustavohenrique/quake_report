SHELL := /bin/bash
go := $(shell which go)

setup: install
install:
	$(go) mod tidy
	$(go) install golang.org/x/tools/cmd/goimports@latest

run:
	@$(go) run main.go

test: tests
tests:
	@$(go) test -v -race -failfast ./...

ci: coverage
coverage:
	@$(go) test -v -race -failfast -p 1 -covermode atomic -coverprofile=profile.cov `go list ./... | grep -v -E -f .testignore` | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''

lint:
	@goimports -w src
