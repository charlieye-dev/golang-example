package main

import (
  "fmt"
)

type Person struct {
  name string
  age int
  address Address
}

type Address struct {
  nationality, city string
}

func main() {
  var person Person
  person.name = "jenny"
  person.age = 18
  person.address = Address {
    nationality : "China",
    city : "Shanghai",
  }

  fmt.Println(person)

  personi := Person {
    "danny", 18, Address {"China", "Beijing"},
  }

  fmt.Println(personi.address.city)
}
