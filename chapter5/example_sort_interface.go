package main

import (
	"fmt"
	"sort"
	"strings"
)

// set type
type CaseInsensitive []string

/**
interface
Len
Less
Swap
*/
func (c CaseInsensitive) Len() int {
	return len(c)
}

func (c CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) || (strings.ToLower(c[i]) == strings.ToLower(c[j])) && c[i] < c[j]
}

func (c CaseInsensitive) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {
	example_case_insensitive_sort()
}

func example_case_insensitive_sort() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})

	sort.Sort(apple)
	fmt.Println(apple)
}
