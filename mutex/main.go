package main

import (
	"fmt"
	"sync"
	"time"
)

var n = 0
var mu = sync.Mutex{}

func count() {
	fmt.Println(n)
	n = n + 1
}

func main() {
	go count()
	go count()
	time.Sleep(1 * time.Hour)
}


/**
package main

import (
	"fmt"
	"sync"
	"time"
)

var n = 0
var mu = sync.Mutex{}

func count() {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(n)
	n = n + 1
}

func main() {
	go count()
	go count()
	time.Sleep(1 * time.Hour)
}
 */