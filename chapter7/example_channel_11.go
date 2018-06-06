package main

import (
	"fmt"
)

// fan out
func main() {
	select {
	case n := <-c1:
		fmt.Println(n, "is from c1")
	case n := <-c2:
		fmt.Println(n, "is from c2")
	case c3 <- f():
		fmt.Println("1 is sent to c3")
	default:
		fmt.Println("No channel is ready")
	}

}
