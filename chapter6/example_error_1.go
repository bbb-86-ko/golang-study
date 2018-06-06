package main

import "fmt"

type ID int

type ErrNegativeID ID

type error interface {
	Error() string
}

func (e ErrNegativeID) Error() string {
	return fmt.Sprintf("ID %d is negative", e)
}

func main() {
	var err error = ErrNegativeID(-1)
	fmt.Println(err)
}
