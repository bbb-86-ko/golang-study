package main

import "fmt"

func main() {
	c := make(chan int, 10)
	c <- 5
	fmt.Print(<-c)
}
