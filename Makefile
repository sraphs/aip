MOD_NAME := github.com/sraphs/go-starter

# Git variables
GIT_COMMIT    = $(shell git rev-parse HEAD)
GIT_SHA       = $(shell git rev-parse --short HEAD)
GIT_TAG       = $(shell git describe --tags --always)
GIT_DIRTY     = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
GIT_REPO_URL  = $(shell git remote get-url origin | sed -e 's|git@\(.*\):\(.*\)\.git|https://\1/\2|g')

# Go variables
GO        ?= go
GOOS      := $(shell $(GO) env GOOS)
GOARCH    := $(shell $(GO) env GOARCH)
GOHOST    := GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO)
GOVERSION := $(shell go version | cut -d ' ' -f 3 | cut -d '.' -f 2)

# Build variables
VERSION  ?= $(GIT_TAG)

LDFLAGS ?= "-s -w -X main.version=$(VERSION)"

.DEFAULT_GOAL := help

###############
##@ Initial

.PHONY: init
init: ## Init environment
	@ $(MAKE) --no-print-directory log-$@
	go install github.com/goreleaser/goreleaser@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin 

.PHONY: rename
rename: ## Rename Go module refactoring
	@ $(MAKE) --no-print-directory log-$@
	@echo "Enter new go module-name:" \
		&& read new_module_name \
		&& echo "new go module-name: '$${new_module_name}'" \
		&& echo -n "Are you sure? [y/N]" \
		&& read ans && [ $${ans:-N} = y ] \
		&& echo -n "Please wait..." \
		&& find . -type f -not -path '*/\.git*' -exec sed -i "s|${MOD_NAME}|$${new_module_name}|g" {} \; \
		&& echo "new go module-name: '$${new_module_name}'!" \
        && git add . && git commit -m "rename go module-name to '$${new_module_name}'"

###############
##@ Development

.PHONY: dev
dev: ## Dev
	@ $(MAKE) --no-print-directory log-$@
	CGO_ENABLED=0 $(GOHOST) run -ldflags=$(LDFLAGS) ./...

.PHONY: clean
clean: ## Clean workspace
	@ $(MAKE) --no-print-directory log-$@
	go clean

.PHONY: check 
check: test lint ## Run tests and linters

.PHONY: test
test: ## Run tests
	@ $(MAKE) --no-print-directory log-$@
	@ $(MAKE) --no-print-directory log-$@
	go test -race -coverprofile coverage.out -covermode=atomic ./...
	go tool cover -func=coverage.out

.PHONY: lint
lint: ## Run golint linter
	@ $(MAKE) --no-print-directory log-$@
	golangci-lint run --fast

#########
##@ Build

.PHONY: build
build: ## Build
	@ $(MAKE) --no-print-directory log-$@
	CGO_ENABLED=0 $(GOHOST) build -ldflags=$(LDFLAGS) ./...

########
##@ Help
.PHONY: info
info: ## Display build info
	@ $(MAKE) --no-print-directory log-$@
	@echo "Mod Name:           ${MOD_NAME}"
	@echo "Version:            ${VERSION}"
	@echo ""
	@echo "Git Tag:            ${GIT_TAG}"
	@echo "Git Commit:         ${GIT_COMMIT}"
	@echo "Git Tree State:     ${GIT_DIRTY}"
	@echo "Git Repository URL: ${GIT_REPO_URL}"
	@echo ""
	@echo "Go  Version:        ${GOVERSION}"
	@echo "GO  OS:             ${GOOS}"
	@echo "GO  ARCH:           ${GOARCH}"

.PHONY: help
help:   ## Display this help 
	@awk -v "col=\033[36m" -v "nocol=\033[0m" ' BEGIN { FS = ":.*##" ; printf "Usage:\n  make %s<target>%s\n", col, nocol } /^[a-zA-Z_-]+:.*?##/ { printf "  %s%-12s%s %s\n", col, $$1, nocol, $$2 } /^##@/ { printf "\n%s%s%s\n", nocol, substr($$0, 5), nocol }' $(MAKEFILE_LIST)

log-%: 
	@grep -h -E '^$*:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN { FS = ":.*?## " }; { printf "\033[36m==> %s\033[0m\n", $$2}'
