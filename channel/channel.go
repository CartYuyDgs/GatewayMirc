package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	fmt.Println(c)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println(<-c)
		wg.Done()
	}()
	c <- 1
	wg.Wait()
	close(c)
	fmt.Println(c)
}
