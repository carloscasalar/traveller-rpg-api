# Traveller RPG API

This is an API for the Traveller RPG. It is a Cloudflare Worker that uses Go to run the server. 
It is based on the [Cloudflare Worker Go template](https://github.com/syumai/workers)

Why with GO? Just for fun :)

## Notice

- A free plan Cloudflare Workers only accepts ~1MB sized workers.
  - Go Wasm binaries easily exceeds this limit, so **you'll need to use a paid plan of Cloudflare Workers** (which accepts ~5MB sized workers).

## Usage

- `main.go` includes simple HTTP server implementation. 

## Requirements

- Node.js
- [wrangler](https://developers.cloudflare.com/workers/wrangler/)
  - just run `npm install -g wrangler`
- Go 1.21.0 or later

## Development

### Commands

```
make dev     # run dev server
make build   # build Go Wasm binary
make deploy # deploy worker
```

### Testing dev server

- Just send HTTP request using some tools like curl.

```
$ curl http://localhost:8787/hello
Hello!
```

```
$ curl -X POST -d "test message" http://localhost:8787/echo
test message
```
