package main

import "fmt"

func find(flag bool, nums ...int) {
  if flag == true {
    fmt.Println("Hello")
  } else {
    return
  }

  fmt.Printf("nums's type is %T\n", nums)
  for _, v := range(nums) {
    fmt.Println(v)
  }
  fmt.Println("")
}

func main() {
  nums := []int{1, 2, 3}
  find(true, nums...)
  find(false, 1, 2, 3, 4)
  find(true, 1, 2, 3, 4, 5)
}
