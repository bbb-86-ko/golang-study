package main

import "fmt"

func main()  {
  fmt.Println(hasDupeRune_1("숨바꼭질."))
  fmt.Println(hasDupeRune_2("다시합시다."))
}

func hasDupeRune_1(s string) bool  {
  runeSet := map[rune]bool{}
  for _, r:= range s{
    if runeSet[r]{
      return true
    }
    runeSet[r] = true
  }
  return false
}


func hasDupeRune_2(s string) bool  {
  runeSet := map[rune]struct{}{}
  for _, r:= range s{
    if _, exists := runeSet[r]; exists {
      return true
    }
    runeSet[r] = struct{}{}
  }
  return false
}
