package main

import "fmt"

func main() {
  a := [6]int{1, 2, 3, 4, 5, 6}

  fmt.Println(a)
  for i, v := range(a) {
    fmt.Println(i, v)
  }

  fmt.Println(" ")

  // slice is a reference of arrary
  var b []int = a[2:5]
  for i, v := range(b) {
    fmt.Println(i, v)
  }
  b = append(b, 8, 9)
  fmt.Println(b)
  fmt.Println("slice b len is", len(b), ", cap is", cap(b))
  b = b[:cap(b)]
  fmt.Println(b)
  fmt.Println("slice b len is", len(b), ", cap is", cap(b))

  var c []int = a[:]
  fmt.Println(c)

  fmt.Println("")

  d := make([]int, 5, 5)
  for i := range(d) {
    d[i] = i
  }
  fmt.Println(d)
  fmt.Println("d slice len is", len(d), "cap is", cap(d))
  d = append(d, 5)
  fmt.Println(d)
  fmt.Println("d slice len is", len(d), "cap is", cap(d))

  fmt.Println("")

  var e []int
  if e == nil {
    fmt.Println("nil slice e")
    e = append(e, 1, 2, 3, 4, 5)
    e = append(e, b...)
    fmt.Println(e)
    fmt.Println("e slice len is", len(e), ", cap is", cap(e))
  }

  var f []int = a[2:5]
  g := make([]int, len(f), len(f))
  copy(g, f)
  g[2] = 20
  fmt.Println(g)
}
