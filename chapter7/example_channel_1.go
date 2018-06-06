package main

func main() {
	c1 := make(chan int)
	var chan int c2 = c1
	var <-chan int c3 = c1		// recive only (like consum??)
	var chan<- int c4 = c4		// send only
}
