package main

import (
	"fmt"
	"sync"
)

var n = 0
var mu sync.Mutex
var cond = sync.NewCond(&mu)

func count() {
	for {
		mu.Lock()
		if n == 10000 {
			mu.Unlock()
			break
		} else {
			n++
		}
		mu.Unlock()
	}
}

func main()  {
	go count()
	go count()

	for {
		mu.Lock()
		fmt.Println("waiting...")
		if n == 10000 {
			fmt.Println("finish")
			mu.Unlock()
			break
		}
		mu.Unlock()
	}
}


/**
package main

import (
	"fmt"
	"sync"
)

var n = 0
var mu sync.Mutex
var cond = sync.NewCond(&mu)

func count() {
	for {
		mu.Lock()
		if n == 10000 {
			cond.Broadcast()
			mu.Unlock()
			break
		} else {
			n++
		}
		mu.Unlock()
	}
}

func main()  {
	go count()
	go count()

	for {
		mu.Lock()
		fmt.Println("waiting...")
		if n != 10000 {
			cond.Wait()
		} else {
			fmt.Println("finish")
			break
		}
		mu.Unlock()
	}
}

 */