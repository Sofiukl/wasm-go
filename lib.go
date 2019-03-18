package main

import (
	"syscall/js"
)

func wrapHTML(args []js.Value) {
	s := "<h1>" + args[0].String() + "</h1>"
	println(s)
}
func welcome(args []js.Value) {
	js.Global().Get("document").
		Call("getElementById", "welcome-msg").
		Set("innerHTML", "Welcome to WASM.. New way to develop app")
}
func registerCallbacks() {
	js.Global().Set("generateHTML", js.NewCallback(wrapHTML))
	js.Global().Set("welcome", js.NewCallback(welcome))
}

func main() {
	c := make(chan struct{}, 0)
	println("Hello, WebAssembly using Go!")
	registerCallbacks()
	<-c
}
