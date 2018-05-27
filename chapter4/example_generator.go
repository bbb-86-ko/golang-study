package main

import "fmt"

func main() {
	example_new_int_generator()
}

func new_int_generator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

func example_new_int_generator() {
	gen1 := new_int_generator()
	gen2 := new_int_generator()
	fmt.Println(gen1(), gen1(), gen1(), gen1())
	fmt.Println(gen2(), gen2(), gen2(), gen2())
	fmt.Println(gen1(), gen1(), gen1(), gen1())
}
