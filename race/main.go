package main

import "fmt"

var (
	counter = 0
)

func updateCounter() {
	counter++
}

func main() {
	go updateCounter()
	fmt.Println(counter)
}
