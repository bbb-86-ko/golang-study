package main

import (
	"fmt"
	"time"
)

type status int
type Deadline time.Time

const (
	UNKNOWN status = iota * 10
	TODO
	DONE
)

type Task struct {
	title    string
	status   status
	Deadline *Deadline
}

func main() {
	d1 := Deadline(time.Now().Add(-4 * time.Hour))
	d2 := Deadline(time.Now().Add(4 * time.Hour))
	t1 := Task{"4h ago", TODO, &d1}
	t2 := Task{"4h later", TODO, &d2}
	t3 := Task{"no due", TODO, nil}
	fmt.Println(t1.OverDue())
	fmt.Println(t2.OverDue())
	fmt.Println(t3.OverDue())

}

func (d *Deadline) OverDue() bool {
	return d != nil && time.Time(*d).Before(time.Now())
}

func (t Task) OverDue() bool {
	return t.Deadline.OverDue()
}
