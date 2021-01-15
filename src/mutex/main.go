package main

import (
	"fmt"
	"sync"
)

var x int

func add(w *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	w.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go add(&w, &m)
	}
	w.Wait()

	fmt.Println(x)
}
