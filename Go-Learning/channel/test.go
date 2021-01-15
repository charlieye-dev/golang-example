package main

import (
  "time"
  "fmt"
)

var a int

func Count(ch chan int) {
    a = a + 1
    ch <- a
    fmt.Println("Counting")
}

func main() {
    chs := make([]chan int, 10)

    for i := 0; i < 10; i++ {
        chs[i] = make(chan int)
        go Count(chs[i])
    }

    for _, ch := range(chs) {
        <-ch
    }
    time.Sleep(time.Second)
}
