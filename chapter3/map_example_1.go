package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func main() {
	// var m map[keyType]valueType
	// m = make(map[keyType]valueType)

	// m := make(map[keyType]valueType)

	// m := map[keyType]valueType{}

	example_count_1()
	fmt.Println()
	example_count_2()
	fmt.Println()
}

func count(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}

func TestCount1(t *testing.T) {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	if !reflect.DeepEqual(map[rune]int{'가': 1, '나': 2, '다': 3}, codeCount) {
		t.Error("codeCount mismatch:", codeCount)
	}
}

func TestCount2(t *testing.T) {
	// default map
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	if len(codeCount) != 3 {
		t.Error("codeCount:", codeCount)
		t.Fatal("count should be 3 but:", len(codeCount))
	}
	if codeCount['가'] != 1 || codeCount['나'] != 2 || codeCount['다'] != 1 {
		t.Error("codeCount mismatch:", codeCount)
	}
}

func example_count_1() {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	for _, key := range []rune{'가', '나', '다'} {
		fmt.Println(string(key), codeCount[key])
	}
}

func example_count_2() {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	var keys sort.IntSlice
	for key := range codeCount {
		keys = append(keys, int(key))
	}
	sort.Sort(keys)
	for _, key := range keys {
		fmt.Println(string(key), codeCount[rune(key)])
	}
}
