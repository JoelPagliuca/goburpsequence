NAME := goburpsequence
VERSION := $(shell cat VERSION.txt)

.PHONY: help
help: ## print this message and exit
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: $(NAME) ## Builds a dynamic executable or package

$(NAME): $(wildcard *.go) VERSION.txt
	@echo "+ $@"
	go build -o $(NAME) .