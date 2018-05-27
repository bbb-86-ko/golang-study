package main

import "fmt"

func main() {
	example_literal_1()
	func() {
		fmt.Println("hello1")
	}()
	printHello := func() {
		fmt.Println("hello2")
	}
	printHello()
}

func example_literal_1() {
	func() {
		fmt.Println("hello0")
	}()
}

func add(a, b int) int {
	return a + b
}
