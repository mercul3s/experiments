package main

import (
	"fmt"
	"time"

	"github.com/mercul3s/timer_example/timer"
)

type event struct {
	ttl     time.Duration
	message string
}

func main() {
	e1 := event{
		ttl:     5,
		message: "hello world",
	}

	e1Timer := timer.GetTimer(e1.ttl)
	<-e1Timer.C
	fmt.Println(e1.message)
}
