package main

import (
  "fmt"
  "time"
)

func printNum(ch chan int) {
  for i := 0; i < 10; i++ {
    ch <- i
    fmt.Println("input", i)
  }
  close(ch)
}

func main() {
  ch := make(chan int, 2)
  go printNum(ch)

  for v := range ch {
    time.Sleep(3 * time.Second)
    fmt.Println(v)
  }
}
