/*
 * Type convert
 *
 * go is a strongly type language, it can not convert type automatically.
 * it can convert variable type by T(v).
 */

package main

import "fmt"

func main() {
  x := 3
  y := 3.5
  sum := x + int(y)

  fmt.Println("sum is", sum)
}
