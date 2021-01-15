package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Printf("Runtime os: ")
  switch os := runtime.GOOS; os {
  case "drawwin":
    fmt.Println("Window");
  case "linux":
    fmt.Println("Linux");
  default:
    fmt.Println(os);
  }
}
