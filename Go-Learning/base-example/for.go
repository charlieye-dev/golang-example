package main

import "fmt"

func main() {
  sum := 1

  for i := 0; i < 10; i++ {
    sum += sum
  }

  fmt.Println(sum);

  for sum != 0 {
    fmt.Println("true");
    sum = 0
  }
}
