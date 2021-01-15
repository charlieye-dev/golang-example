package main

/*
#include <stdio.h>

static const char* hello() {
  const char* p = "hello world";
  return p;
}
*/
import "C"
import "fmt"

func main() {
  err := C.GoString(C.hello())
  fmt.Println(err)
}
