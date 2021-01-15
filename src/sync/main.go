package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
)

func worker(id int) {

	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	<-ctx.Done()
	fmt.Printf("Worker %d done\n", id)
}

func initialization() {
	hello := new(sync.WaitGroup)
	wg = hello
}

func main() {

	initialization()

	ctx, cancel = context.WithCancel(context.Background())

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i)
		time.Sleep(time.Second)
	}

	cancel()
	wg.Wait()
}
