package main

import "fmt"

func main() {
	// i : byte 위치 : int
	// r : rune type : int32 alias
	for i, r := range "A B C" {
		fmt.Println(i, r)
	}
	fmt.Println(len("A B C"))

	for _, r := range "A B C" {
		fmt.Println(string(r), r)
	}
	fmt.Println(len("A B C"))

}
