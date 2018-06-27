package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 3)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
