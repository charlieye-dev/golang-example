package main

import (
	"fmt"
	"sync"
)

var x int

func add(w *sync.WaitGroup, c chan bool) {
	c <- true
	x = x + 1
	<-c
	w.Done()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go add(&w, ch)
	}
	w.Wait()

	fmt.Println(x)
}
