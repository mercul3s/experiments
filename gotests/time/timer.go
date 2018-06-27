package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Add(500 * time.Millisecond)
	fmt.Println(t.Before(time.Now().Add(750 * time.Millisecond)))
	fmt.Println(time.Now().After(t))
	time.Sleep(750 * time.Millisecond)
	fmt.Println(time.Now().After(t))
}
