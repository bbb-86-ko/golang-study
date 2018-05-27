package main

import "fmt"

type BinOp func(int, int) int
type BinSub func(int, int) int

func main() {
	OpTheeAndFour(func(a, b int) int {
		return a + b
	})

	sub := BinOpToBinSub(func(a, b int) int {
		return a + b
	})

	sub(5, 7)
	sub(5, 7)

	count := sub(5, 7)
	fmt.Println("count : ", count)
}

func OpTheeAndFour(f BinOp) {
	fmt.Println(f(3, 4))
}

func BinOpToBinSub(f BinSub) BinSub {
	var count int
	return func(a, b int) int {
		fmt.Println(f(a, b))
		count++
		return count
	}
}
