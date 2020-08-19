package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var n int64

func genId() {
	for {
		id := atomic.AddInt64(&n, 1)
		if id % 100000000 == 0 {
			fmt.Println(id)
		}
	}
}

func main()  {
	go genId()
	go genId()

	time.Sleep(1 * time.Hour)
}
