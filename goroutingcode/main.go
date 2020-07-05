package main

import (
	"fmt"
	"time"
)

func action() {
	fmt.Println("Test Grouting")
}

func main() {
	go action()
	time.Sleep(2)
}
