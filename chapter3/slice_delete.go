package main

import "fmt"

func main() {
	example_slice_delete_1()
	fmt.Println()
	example_slice_delete_2()
	fmt.Println()

}

func example_slice_delete_1() {
	i := 2
	nums := []int{30, 20, 50, 10, 40}
	fmt.Println(nums)
	nums = append(nums[:i], nums[i+1:]...)
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
}

func example_slice_delete_2() {
	i := 2
	nums := []int{30, 20, 50, 10, 40}
	fmt.Println(nums)
	nums[i] = nums[len(nums)-1]
	fmt.Println(nums)
	nums = nums[:len(nums)-1]
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
}
