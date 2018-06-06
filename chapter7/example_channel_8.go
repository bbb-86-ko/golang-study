package main

import "fmt"

type IntPipe func(<-chan int) <-chan int

func main() {

	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 6
		c <- 7
	}()

	for num := range PlusOne(PlusOne(c)) {
		fmt.Println(num)
	}

	c2 := make(chan int)
	go func() {
		defer close(c2)
		c2 <- 10
		c2 <- 20
		c2 <- 30
	}()

	for num := range Chain(PlusOne, PlusOne, PlusOne, PlusOne)(c2) {
		fmt.Println(num)
	}
}

func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}
