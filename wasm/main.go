package main

import (
	"syscall/js"
	"time"
)

func fib(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func fibFunc(this js.Value, args []js.Value) interface{} {
	callback := args[len(args)-1]

	go func() {
		time.Sleep(3 * time.Second)
		v := fib(args[0].Int())
		callback.Invoke(v)
	}()

	js.Global().Get("ans").Set("innerHTML", "waiting 3s...")
	// v := numEle.Get("value")
	// if num, err := strconv.Atoi(v.String()); err == nil {
	// 	ansEle.Set("innerHTML", js.ValueOf(fib(num)))
	// }
	return nil
}

var (
	document = js.Global().Get("document")
	numEle   = document.Call("getElementById", "num")
	ansEle   = document.Call("getElementById", "ans")
	btnEle   = document.Call("getElementById", "btn")
)

func main() {
	// alert := js.Global().Get("alert")
	// alert.Invoke("hello world")
	done := make(chan int, 0)
	// btnEle.Call("addEventListener", "click", js.FuncOf(fibFunc))
	js.Global().Set("fibFunc", js.FuncOf(fibFunc))
	<-done
}
