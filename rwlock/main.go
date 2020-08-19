package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMu = sync.RWMutex{}

func printer(name string) {
	for {
		rwMu.RLock()
		fmt.Println(name, "sleep")
		time.Sleep(5 * time.Second)
		fmt.Println(name, "awake")
		rwMu.RUnlock()
	}
}

func main() {
	go printer("a")
	go printer("b")

	for {
		rwMu.Lock()
		fmt.Println("write")
		rwMu.Unlock()
	}
}
