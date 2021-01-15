package main

import "fmt"
import "basic/area"

func caculate(price, num int) (sum, number int) {
  sum = price * num
  return
}

func init() {
  fmt.Println("I'm main package")
  area.Area()
}

func main() {
  sum, _ := caculate(10, 20)
  fmt.Println("sum is", sum)
  fmt.Println("Hello,", area.Str)
}
