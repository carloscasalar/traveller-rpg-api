.PHONY: dev
dev:
	wrangler dev

.PHONY: build
build:
	go run github.com/syumai/workers/cmd/workers-assets-gen@v0.23.1 -mode=go
	GOOS=js GOARCH=wasm go build -o ./build/app.wasm .

.PHONY: deploy
deploy:
	wrangler deploy

# Install the required tools for go generators
install-tools:
	@echo "Parsing tools.go and installing dependencies..."
	@cd tools && go list -e -f '{{join .Imports " "}}' tools.go | xargs -t -n 1 $(GO_BIN) install
	@echo "all tools installed"