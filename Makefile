PACKAGES := $(shell go list ./...)
name := $(shell basename ${PWD})

# Go parameters
GO := go
GOBUILD := $(GO) build
GOTEST := $(GO) test
GOCLEAN := $(GO) clean

# Main package directory
CMD_DIR := ./cmd
MAIN_PACKAGE := main.go

# Output binary name
BINARY_NAME := rari

all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## init: initialize project (make init module=github.com/user/project)
.PHONY: init
init:
	go mod init ${module}
	go install github.com/cosmtrek/air@latest
	asdf reshim golang

## vet: vet code
.PHONY: vet
vet:
	go vet $(PACKAGES)

## test: run unit tests
# .PHONY: test
# test:
# 	go test -race -cover $(PACKAGES)

# ## build: build a binary
# .PHONY: build
# build: test
# 	go build -o cmd/main.go -v

# Test the code
test:
	$(GOTEST) -v ./...

# Build the executable
build: test
	$(GOBUILD) -o $(BINARY_NAME) $(CMD_DIR)/$(MAIN_PACKAGE)

# Run the application
run: build
	air

## docker-build: build project into a docker container image
.PHONY: docker-build
docker-build: test
	GOPROXY=direct docker buildx build -t ${name} .

## docker-run: run project in a container
.PHONY: docker-run
docker-run:
	docker run -it --rm -p 8080:8080 ${name}

## css: build tailwindcss
.PHONY: css
css:
	tailwindcss -i css/input.css -o css/output.css --minify

## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	tailwindcss -i css/input.css -o css/output.css --watch
