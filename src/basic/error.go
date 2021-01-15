package main

import "fmt"
import "errors"

func hello() error {
  return errors.New("hello")
}

func main() {
  err := hello()

  fmt.Println(err)
}
