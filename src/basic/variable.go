/*
 * Variables
 *
 * 1. short hand delaration can only used as local variable.
 * 2. at least one new variable should be provided on left side of ":=".
 * 3. local variable must be used after declaring.
 */

package main

import "fmt"

var(
  k = 5
  l string
)

func main() {
	var a int
	fmt.Println("not initialvalue a", a)

  var b int = 2
	fmt.Println("initialvalue b", b)

  var c, d int = 3, 4
	fmt.Println("mutiple value c", c, ", d", d)

  var (
    e int = 5
    f string = "hello"
    g int

    // following is ok, type reference
    /*
    e = 5
    f = "hello"
    g int
    */
  )
  g = 6
  fmt.Println("multiple type variable: e is", e, "f is", f, "g is", g)

  // short hand declaration can only be used as local variable
  h, i := 7, 8
  i, j := 9, 10
  // following is unvalid because no new variables on left side of ":="
   i, j := 11, 12
  fmt.Println("short hand declaration h", h, "i", i, "j", j)
}

