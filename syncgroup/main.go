package main

import (
	"fmt"
	"sync"
)

var syncGroup = sync.WaitGroup{}

func worker(n int) {
	fmt.Println(n)
	syncGroup.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		syncGroup.Add(1)
		go worker(i)
	}
	syncGroup.Wait()
	fmt.Println("all done")
}
