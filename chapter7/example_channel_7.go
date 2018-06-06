package main

import "fmt"

func main() {
	c1 := make(chan int, 10)
	c1 <- 10
	val1, ok1 := <-c1
	fmt.Println(val1, ok1)

	c2 := make(chan int, 10)
	close(c2)
	val2, ok2 := <-c2
	fmt.Println(val2, ok2)
}
