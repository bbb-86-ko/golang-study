package main

import "fmt"

func main() {
	example_slice_insert_1()
	fmt.Println()
	example_slice_insert_2()
	fmt.Println()
	example_slice_insert_3()
	fmt.Println()

}

func example_slice_insert_1() {
	nums := []int{30, 20, 50, 10, 40}
	fmt.Println(nums)
	if i := 2; i < len(nums) {
		nums = append(nums[:i+1], nums[i:]...)
		fmt.Println(nums)
		nums[i] = 3000
	} else {
		fmt.Println(nums)
		nums = append(nums, 2000)
	}

	fmt.Println(nums)
}

func example_slice_insert_2() {
	i := 2
	nums := []int{30, 20, 50, 10, 40}
	nums = append(nums, 2000)
	fmt.Println(nums)
	copy(nums[i+1:], nums[i:])
	fmt.Println(nums)
	nums[i] = 2000
	fmt.Println(nums)
}

func example_slice_insert_3() {
	i := 2
	nums := []int{30, 20, 50, 10, 40}
	x := []int{300, 200, 500}
	fmt.Println(nums)
	nums = append(nums, x...)
	fmt.Println(nums)
	fmt.Println(x)
	copy(nums[i+len(x):], nums[i:])
	fmt.Println(nums)
	copy(nums[i:], x)
	fmt.Println(nums)
}
