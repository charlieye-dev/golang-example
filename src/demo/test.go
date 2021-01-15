package main

import (
    "fmt"
)

type Hello struct {
  a string
  b string
}

func main() {
  hello := map[Hello]int{
    Hello{a:"a", b:"b",}: 1,
    Hello{a:"c", b:"d",}: 2,
  }

  for key, value := range hello {
   fmt.Println(key, value)
  }

  a := Hello{
    a: "a",
    b: "b",
  }

  b := Hello{
    a: "a",
    b: "c",
  }

  if key, ok := hello[a]; !ok {
    fmt.Println("xxx")
    fmt.Println(key, ok)
  }

  if key, ok := hello[b]; !ok {
    fmt.Println("yyy")
    fmt.Println(key, ok)
  }

  d := true

  for d {
    fmt.Println("---")
  }
}
