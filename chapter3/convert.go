package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	var k int64
	var f float64
	var s string
	var err error

	i, err = strconv.Atoi("350")
	k, err = strconv.ParseInt("cc7fdd", 16, 32)
	k, err = strconv.ParseInt("0xcc7fdd", 0, 32)
	f, err = strconv.ParseFloat("3.14", 64)
	s = strconv.Itoa(340)
	s = strconv.FormatInt(13402077, 16)

	fmt.Println(i, k, f, s, err)

}
