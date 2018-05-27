package main

import (
	"fmt"
	"time"
)

type status int

const (
	UNKNOWN status = iota * 10
	TODO
	DONE
)

type Task struct {
	title  string
	status status
	due    *time.Time
}

func main() {
	var myTask2 = Task{"laundry", UNKNOWN, nil}
	var myTask3 = Task{"laundry", TODO, nil}
	var myTask4 = Task{"laundry", DONE, nil}
	fmt.Println(myTask2)
	fmt.Println(myTask3)
	fmt.Println(myTask4)
}
