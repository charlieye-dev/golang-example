package main

import "fmt"

func hello(done chan<- string) {
  done <- "hello, world"
}

func main() {
  done := make(chan string)
//  go hello(done)
  done <- "hello"

  fmt.Println("hello")
}
