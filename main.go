package main

import (
	"sync"
	"fmt"
)

func main() {
	fmt.Println("hello")
	fmt.Println("bye")
}

func mult(a, b int) int {
	var lock sync.Mutex

	l := lock
	l.Lock()
	l.Unlock()

	return a * b
}
