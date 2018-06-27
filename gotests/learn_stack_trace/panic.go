package main

import "fmt"

type trace struct{}

func main() {
	slice := make([]string, 2, 4)
	
	var t trace
	t.Example(slice, "hello", 10)
}

func (t *trace) Example(slice []string, str string, i int) {
	fmt.Printf("Receiver Address: %p\n", t)
	panic("Want stack trace")
}
