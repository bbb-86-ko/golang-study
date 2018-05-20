package main

import "fmt"

func main() {
	var x int  //	자료형 선언
	var y = 10 //	자료형 추론
	z := 20    // 자료형 추론
	fmt.Println("Hello, PlayGround", x, y, z)

	var a = 10
	b := "little"
	b += " Gophers"
	fmt.Println("Hello, PlayGround", a, b)

}

// func(int, int)
// func(int)	int
// func(int, func(int, int)) func(int) int
