package main

import (
	"fmt"
	"sync"
)

func action() {
	fmt.Println("Test Grouting")
}

func main() {
	//go action()
	//time.Sleep(2)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			fmt.Printf("go routing %d\n", num)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
