package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	outTime := time.Now().Add(time.Millisecond * 50)
	ctx, cancel := context.WithDeadline(context.Background(), outTime)
	defer cancel()
LOOP:
	for {
		select {
		case <-time.After(time.Millisecond * 10):
			m := <-time.After(time.Millisecond * 10)
			fmt.Println(m)
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break LOOP
		}
	}
}
