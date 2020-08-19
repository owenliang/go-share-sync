package main

import "fmt"

var n = 0

func main() {
	fmt.Println(n)
}

func init() {
	n = 1
}