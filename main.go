package main

import (
    "fmt"
    "syscall/js"
)

func main() {
    fmt.Println("Go on Wasm!")

    js.Global().Call("updateDOM", "Go => Wasm")
}
