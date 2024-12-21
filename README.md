# Traveller RPG API

This is an API for the Traveller RPG. It is a Cloudflare Worker that uses Go to run the server. 
It is based on the _awesome_ [Cloudflare Worker Go template](https://github.com/syumai/workers)

Why with GO? Just for fun :)

## Notice

- A free plan Cloudflare Workers only accepts ~3MB sized workers. Using regular Go Wasm binaries easily exceeds this  
limit, but building with TinyGo can reduce the size, in this case from around 5MB to around 950KB.

## Usage

- `main.go` includes simple HTTP server implementation. 

## Requirements

- Node.js
- [wrangler](https://developers.cloudflare.com/workers/wrangler/)
  - just run `npm install -g wrangler`
- Go 1.23.1 or later
- [TinyGo](https://tinygo.org/) to build Go Wasm binary, you need to install TinyGo.

## Development

### Commands

```
make dev     # run dev server
make build   # build Go Wasm binary
make deploy # deploy worker
```

### Testing dev server

- Just send HTTP request using some tools like curl.

```bash
$ curl -X POST http://localhost:8787/api/npcs \
-H "Content-Type: application/json" \
-d '{
  "role": "pilot",
  "citizen_category": "average",
  "experience": "regular",
  "gender": "unspecified"
}'
```
output
```json
{
  "firsts_name": "Ocean",
  "surname": "Smith",
  "role": "pilot",
  "citizen_category": "average",
  "experience": "regular",
  "characteristics": {
    "DEX": 9,
    "EDU": 7,
    "END": 5,
    "INT": 8,
    "SOC": 6,
    "STR": 7
  },
  "skills": [
    "Pilot (Spacecraft)-2",
    "Astrogation-2",
    "Electronic (Sensors)-1",
    "Gunnery-1",
    "Mechanic-0",
    "Leadership-0",
    "Vacc Suit-0",
    "Electronics-0",
    "Drive-0"
  ]
}
```
