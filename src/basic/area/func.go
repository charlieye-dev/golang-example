package area

import (
  "fmt"
  "log"
)

const Str = "China!"
var a int = -7

func init() {
  fmt.Println("I'm area package")
  if a < 0 {
    log.Fatal("Value less than zero!!!")
  }
}

func Area() {
  fmt.Println("Hello, world!")
}
