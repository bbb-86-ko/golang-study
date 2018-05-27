package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	b, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))

	deserialContant := Contact{}
	err_2 := json.Unmarshal(b, &deserialContant)
	if err_2 != nil {
		log.Println(err_2)
		return
	}
	fmt.Println(deserialContant)
}
