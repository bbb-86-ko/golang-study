package main

import "fmt"

func main() {
	for fi := range Fibonacci(16) {
		fmt.Println(fi)
	}

	fib := FibonacciGenerator(16)
	for n := fib(); n >= 0; n = fib() {
		fmt.Print(n, ",")
	}
}

func Fibonacci(max int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)

		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()

	return c
}

func FibonacciGenerator(max int) func() int {
	next, a, b := 0, 0, 1
	return func() int {
		next, a, b = a, b, a+b
		if next > max {
			return -1
		}
		return next
	}
}
