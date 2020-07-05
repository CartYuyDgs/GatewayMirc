package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	parent := context.Background()

	ctx, cancel := context.WithCancel(parent)

	runtimes := 0

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Grouting Done")
				return
			default:
				fmt.Printf("Grouting Running Times :%d\n", runtimes)
				runtimes = runtimes + 1

			}

			if runtimes > 5 {
				cancel()
				wg.Done()
			}
		}
	}(ctx)
	wg.Wait()
}
