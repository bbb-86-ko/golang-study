package main

import (
	"fmt"
	"sync"
)

func Min(a []int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, e := range a[1:] {
		fmt.Print("e :", e, " | ")
		if min > e {
			min = e
		}
	}
	fmt.Println()
	return min
}

func ParallelMin(a []int, n int) int {
	if len(a) == 0 {
		return Min(a)
	}
	mins := make([]int, n)
	size := (len(a) + n - 1) / n
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Print("i : ", i, "  ")
			defer wg.Done()
			begin, end := i*size, (i+1)*size
			if end > len(a) {
				end = len(a)
			}
			mins[i] = Min(a[begin:end])
		}(i)
	}
	fmt.Println()
	wg.Wait()

	fmt.Println("mins : ", mins)
	return Min(mins)
}

func main() {
	fmt.Println(Min([]int{
		83, 46, 49, 23, 92, 48, 39, 91, 44, 99, 25, 42, 74, 56, 23,
	}))
	fmt.Println(ParallelMin([]int{
		83, 46, 49, 23, 92, 48, 39, 91, 44, 99, 25, 42, 74, 56, 23,
	}, 5))
}
