NAME := goburpsequence
VERSION := $(shell cat VERSION.txt)

.PHONY: help
help: ## Print this message and exit
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: $(NAME) ## Builds a dynamic executable or package

$(NAME): $(wildcard *.go) VERSION.txt
	@echo "+ $@"
	go build -o $(NAME) .

.PHONY: test
test: ## Runs the go tests
	@echo "+ $@"
	@go test -v $(shell go list ./... | grep -v vendor)

.PHONY: vet
vet: ## Verifies `go vet` passes
	@echo "+ $@"
	@go vet $(shell go list ./... | grep -v vendor)