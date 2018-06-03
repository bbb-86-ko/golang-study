package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("In goroutine")
	}()
	fmt.Println("In main routine")
}
