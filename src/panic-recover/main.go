package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("recover: %v\n", e)
			}
		}()

		close(ch)
		ch <- struct{}{}
	}()

	close(ch)
	<-ch

	time.Sleep(time.Second)

	fmt.Println("vim-go")
}
