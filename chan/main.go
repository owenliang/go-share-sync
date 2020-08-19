package main

import "fmt"

var taskChan = make(chan int)
var exitChan = make(chan int)

func worker() {
	for {
		job, ok := <- taskChan
		if !ok {
			exitChan <- 1
			return
		}
		fmt.Println(job)
	}
}

func main() {
	go worker()
	go worker()

	for i := 0; i < 10; i++ {
		taskChan <- i
	}
	close(taskChan)

	for i := 0; i < 2; i++ {
		<- exitChan
	}
	fmt.Println("all done")
}
