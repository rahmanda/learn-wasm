# Learn WASM

## Prequisites

- Install [tinygo](https://tinygo.org/getting-started/install/linux/)
- Install [emscripten](https://emscripten.org/docs/getting_started/index.html)

## How to Compile

### C/C++

- `cd` to your chosen directory
- Run `emcc -g main.c -o index.html`. This command will generate `index.wasm`, `index.js` and `index.html` on your current directory.

### Golang

- `cd` to your chosen directory
- Run `tinygo build -o main.wasm -target wasm ./main.go`. This command will generate `main.wasm` on current directory.
- Run `cp $(tinygo env TINYGOROOT)/targets/wasm_exec.js .`. This command will generate `wasm_exec.js` on current directory.

## How to Preview Demo

- Run `node ./server.js`
- Open http://localhost:8000/index.html on your browser
