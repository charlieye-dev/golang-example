package main

import "fmt"
import "time"

func test(quitCh chan struct{}) {
  for {
    fmt.Println("world")
    select {
    case <-time.After(1 * time.Second):
      fmt.Println("sleep")
      break
    case <-quitCh:
      fmt.Println("hellor")
      return
    }
  }
}

func main() {
  ch := make(chan struct{})

  go test(ch)

  time.Sleep(30 * time.Second)
  ch <- struct{}{}

  time.Sleep(3 * time.Second)
  fmt.Println("222")
}
