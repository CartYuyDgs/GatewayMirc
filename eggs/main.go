package main

import (
	"fmt"
	"sync"
)

func main() {
	eggs := make(chan int, 10)

	for i := 1; i < 10; i++ {
		eggs <- 10
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			select {
			case egg := <-eggs:
				fmt.Printf("people : %d, Get egg :%d\n", num, egg)
			default:
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
