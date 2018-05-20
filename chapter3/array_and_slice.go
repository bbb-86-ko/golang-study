package main

import (
	"fmt"
)

func main() {
	example_array_1()
	fmt.Println("")
	example_array_2()
	fmt.Println("")
	example_slice_1()
	fmt.Println("")
	example_slice_3()
	fmt.Println("")
	example_slice_append_1()
	fmt.Println("")
	example_slice_append_2()
	fmt.Println("")
	example_slice_cap_1()
	fmt.Println("")
	example_slice_cap_2()
	fmt.Println("")
}

func example_array_1() {
	fruits := [3]string{"사과", "바나나", "토마토"}
	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다.  ", fruit)
	}
}

func example_array_2() {
	fruits := [...]string{"사과", "바나나", "토마토"}
	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다.  ", fruit)
	}
}

func example_slice_1() {
	fruits := []string{"사과", "바나나", "토마토"}
	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다.  ", fruit)
	}
}

func example_slice_2() {
	// var fruits []string
	// fruits = make([]string, 3) // default value set ("" or 0)
	fruits := make([]string, 3)
	fruits[0] = "사과"
	fruits[1] = "바나나"
	fruits[2] = "토마토"
	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다.  ", fruit)
	}
}

func example_slice_3() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])
	fmt.Println(nums[:len(nums)-1])
}

func example_slice_append_1() {
	fruits := []string{"사과", "바나나", "토마토"}
	fmt.Println(fruits)
	fruits = append(fruits, "포도")
	fmt.Println(fruits)
	fruits = append(fruits, "포도", "딸기")
}

func example_slice_append_2() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	fmt.Println(f1)
	fmt.Println(f2)
	f3 := append(f1, f2...)
	fmt.Println(f3)
	f4 := append(f1[:2], f2...)
	fmt.Println(f4)
}

func example_slice_cap_1() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("len:", len(sliced1))
	fmt.Println("cap:", cap(sliced1)) // cap 5
	fmt.Println()

	sliced2 := nums[2:]
	fmt.Println(sliced2)
	fmt.Println("len:", len(sliced2))
	fmt.Println("cap:", cap(sliced2)) // cap 3
	fmt.Println()

	sliced3 := sliced1[:4]
	fmt.Println(sliced3)
	fmt.Println("len:", len(sliced3))
	fmt.Println("cap:", cap(sliced3))
	fmt.Println()

	nums[2] = 100
	fmt.Println(nums, sliced1, sliced2, sliced3)
}

func example_slice_cap_2() {
	// make(type, len[, cap]) len : 3, cap : 5
	nums_1 := make([]int, 3, 5)
	fmt.Println(nums_1)
	fmt.Println("len:", len(nums_1))
	fmt.Println("cap:", cap(nums_1))
	// make(type, len[, cap]) len : 5 cap : 5
	nums_2 := make([]int, 5)
	fmt.Println(nums_2)
	fmt.Println("len:", len(nums_2))
	fmt.Println("cap:", cap(nums_2))
	// make(type, len[, cap]) len : 3 cap : 5
	nums_2 = nums_2[:3]
	fmt.Println(nums_2)
	fmt.Println("len:", len(nums_2))
	fmt.Println("cap:", cap(nums_2))
}
