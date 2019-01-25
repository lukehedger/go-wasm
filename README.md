# Go x Wasm

## Environment
```sh
go version
# >= 1.11.4
```

## Setup
Compile `main.go` to Wasm, output to `main.wasm`:
```sh
GOOS=js GOARCH=wasm go build -o main.wasm
```

- `GOOS` = target operating system, eg. `js` (JavaScript)
- `GOARCH` = architecture, eg. `wasm` (WebAssembly)

Copy the JavaScript Wasm support library provided in the Go source:
```sh
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

## Run
Serve the `index.html` file eg. with [`goexec`](https://github.com/shurcooL/goexec):
```sh
goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
```
