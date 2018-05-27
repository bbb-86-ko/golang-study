package main

import (
	"fmt"
)

type Address struct {
	City  string
	State string
}

type Telephone struct {
	Mobile string
	Direct string
}

type Contact struct {
	Address
	Telephone
}

func main() {
	var c Contact
	c.Mobile = "123-456-789"
	fmt.Println(c.Telephone.Mobile)
	c.Address.City = "San Francisco"
	c.State = "CA"
	c.Direct = "N/A"
	fmt.Println(c)
}
