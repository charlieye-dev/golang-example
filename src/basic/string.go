package main

import "fmt"

func print_byte(name string) {
  for i := 0; i < len(name); i++ {
    fmt.Printf("%x ", name[i])
  }
  fmt.Println("")
}

func print_char(name string) {
  runes := []rune(name)
  for i := 0; i < len(name); i++ {
    fmt.Printf("%c ", runes[i])
  }
  fmt.Println("")
}

func main() {
  name := "hello"
  print_byte(name)
  print_char(name)
}
