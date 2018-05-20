package main

import "fmt"

func main()  {
  example_slice_copy_1()
  fmt.Println()
  example_slice_copy_2()
  fmt.Println()
  example_slice_copy_3()
  fmt.Println()
  example_slice_copy_4()
  fmt.Println()
  example_slice_copy_5()
  fmt.Println()
}

func example_slice_copy_1()  {
  src := []int{30,20,50,10,40}
  dest := make([]int, len(src))
  for idx:= range src{
    dest[idx] = src[idx]
  }
  fmt.Println(src)
  fmt.Println(dest)
}

func example_slice_copy_2()  {
  src := []int{30,20,50,10,40}
  dest := make([]int, len(src))
  copy(dest, src)
  fmt.Println(src)
  fmt.Println(dest)
}

func example_slice_copy_3()  {
  src := []int{30,20,50,10,40}
  dest := make([]int, len(src))
  if n := copy(dest, src); n != len(src) {
    fmt.Println("복사가 덜 됐습니다.")
  }
  fmt.Println(src)
  fmt.Println(dest)
}

func example_slice_copy_4()  {
  src := []int{30,20,50,10,40}
  dest := append([]int(nil), src...)
  fmt.Println(src)
  fmt.Println(dest)
}

func example_slice_copy_5()  {
  src := []int{30,20,50,10,40}
  dest := src               // is not a copy
  fmt.Println(src)
  fmt.Println(dest)
  src[2] = 100
  fmt.Println(src)
  fmt.Println(dest)
}
