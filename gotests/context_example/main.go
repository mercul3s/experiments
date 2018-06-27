package main

import (
	"context"
	"fmt"
	"time"
)

func keepalive(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("in ticker")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}

func main() {
	t := time.Now().Add(5 * time.Second)

	ctx, cancel := context.WithDeadline(context.TODO(), t)
	defer cancel()
	keepalive(ctx)
}
