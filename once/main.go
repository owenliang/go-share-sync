package main

import (
	"fmt"
	"sync"
	"time"
)

var n = 0
var once sync.Once

func initN() {
	n++
}

func worker()  {
	once.Do(initN)
	fmt.Println(n)
}

func main() {
	go worker()
	go worker()

	time.Sleep(1*time.Hour)
}
