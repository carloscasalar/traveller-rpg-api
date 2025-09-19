.PHONY: dev
dev:
	wrangler dev

.PHONY: build
build:
	go tool workers-assets-gen
	tinygo build -o ./build/app.wasm -target wasm -no-debug ./...

.PHONY: deploy
deploy:
	wrangler deploy

# Run the linter
lint:
	@revive -config .revive.toml -formatter friendly ./...

# Run tests
test:
	@go test -v ./...
