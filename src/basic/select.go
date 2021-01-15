package main

import (
  "fmt"
  "time"
)

func hello1(ch chan string) {
//  time.Sleep(3 * time.Second)
  ch <- "Hello"
  close(ch)
}

func hello2(ch chan string) {
//  time.Sleep(4 * time.Second)
  ch <- "World"
  close(ch)
}

func main() {
  ch1 := make(chan string)
  ch2 := make(chan string)
  go hello1(ch1)
  go hello2(ch2)

  for {
    select {
    case s1 := <- ch1:
      fmt.Println(s1)
    case s2 := <- ch2:
      fmt.Println(s2)
    }

    time.Sleep(1 * time.Second)
    fmt.Println("Nice day!!!")
  }
}
