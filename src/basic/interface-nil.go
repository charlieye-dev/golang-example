package main

import "fmt"

type info struct {
  name string
  age int
}

func (i info) describe() {
    fmt.Println("name:", i.name)
    fmt.Println("age:", i.age)
}

func findType(t interface{}) {
  switch i := t.(type) {
  case info:
    i.describe()
  default:
    fmt.Println("Unknow type")
  }
}

func main() {
  i := info{"jenny", 18}
  findType(i)
}
