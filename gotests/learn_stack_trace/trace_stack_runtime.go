package main 

import (
	"runtime/debug"
	"fmt"
)

func main() {
	outer()
}

func outer() {
	inner()
}

func inner () {
	defer func() {
		if err := recover(); err != nil {
			// trace := make([]byte, 1024)
			// count := runtime.Stack(trace, true)
			trace := debug.Stack()

			fmt.Printf("panic in %s", string(trace))
			fmt.Printf("recover from panic: %s\n", err)
			// fmt.Printf("Stack of %d bytes: %s", count, trace)
		}
	}()

	panic("Fake error!")
}

