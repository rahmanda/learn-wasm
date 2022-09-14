// tinygo build -o main.wasm -target wasm ./main.go
package main

func main() {}

//export throwError
func throwError() {
  panic("a problem")
}
