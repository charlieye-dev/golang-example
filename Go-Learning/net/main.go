package main

import "fmt"
import "net"

func main() {
	_, err := net.Listen("tcp", ":22")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("vim-go")
}
