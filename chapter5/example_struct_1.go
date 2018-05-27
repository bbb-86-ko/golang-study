package main

import (
	"fmt"
	"time"
)

var task = struct {
	title string
	done  bool
	due   *time.Time
}{"laundry", false, nil}

type Task struct {
	title string
	done  bool
	due   *time.Time
}

func main() {
	var myTask1 Task
	var myTask2 = Task{"laundry", false, nil}
	var myTask3 = Task{title: "little prince"}
	var myTask4 = Task{title: "little prince", done: true}

	fmt.Println(myTask1)
	fmt.Println(myTask2)
	fmt.Println(myTask3)
	fmt.Println(myTask4)
}
