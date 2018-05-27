package main

import "fmt"

type VerteID int

func main() {
	i := VerteID(100)
	fmt.Println(i)
}

func (id VerteID) String() string {
	return fmt.Sprintf("VertexId(%d)", id)
}
