package main

import "fmt"

func main() {
  var map_str map[string]int

  if map_str == nil {
    fmt.Println(map_str)
    fmt.Println("A nil map")
    map_str = make(map[string]int)
    map_str["one"] = 1
    map_str["two"] = 2
    map_str["three"] = 2
    fmt.Println(map_str)
  }

  _, ok := map_str["one"]
  if ok == true {
    fmt.Println("value is", map_str["one"])
  }

  for key, value := range map_str {
    fmt.Println("key is", key, "value is", value)
  }

  delete(map_str, "one")
  fmt.Println(map_str)

  new_map_str := map_str
  new_map_str["two"] = 20
  fmt.Println(map_str)
}

