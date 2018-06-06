package main

import "fmt"

func main() {
	funcExample_simpleChannel()
}

func funcExample_simpleChannel() {
	c := func() <-chan int {
		c := make(chan int)
		go func() {
			c <- 1
			c <- 2
			c <- 3
			close(c)
		}()
		return c
	}()

	for num := range c {
		fmt.Println(num)
	}
}
