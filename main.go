package main

import (
	"errors"
	"fmt"
)

func A() {
	defer fmt.Println("A")
	/* defer => Defer is a language mechanism that puts your function call into a stack.
	The deferred functions will be executed in reverse order
	when the host function finishes regardless of whether a panic is called or not */

	defer func() {
		if r := recover(); r != nil { // recover => catch the panic
			fmt.Printf("Panic: %+v\n", r)
		}
	}()
	B()
}
func B() {
	defer fmt.Println("B")
	C()
}
func C() {
	defer fmt.Println("C")
	Break()
}
func Break() {
	defer fmt.Println("D")
	panic(errors.New("the show must go on")) // panic => exception
}
func main() {
	A()
}
