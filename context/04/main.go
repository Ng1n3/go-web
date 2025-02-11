package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	time.Sleep(1 * time.Minute)
}

// gen is a broken gnerator that will leak a goroutine.
func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
      select {
      case <-ctx.Done():
        return
      case ch <- n:
        n++
      }
		}
	}()
	return ch
}
