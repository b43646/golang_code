package main

import (
	"D1/ex1/counters"
	"fmt"
)

func main() {
	counter := counters.New(10)
	fmt.Printf("Counter: %d\n", counter)
}
