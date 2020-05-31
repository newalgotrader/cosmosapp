all: install

install: go.sum
		go install -mod=readonly ./cmd/appd
		go install -mod=readonly ./cmd/appcli

go.sum: go.mod
	@echo "--> Checking dependencies"
	go mod verify

test:
	@go test ./...

# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify
