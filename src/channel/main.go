package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{}, 1)
	/*
		timer2 := time.NewTimer(time.Second)
		go func() {
			timer2.Stop()
			fmt.Println("Timer 2 fired")
		}()
	*/
	for {
		fmt.Println(len(ch))
		if len(ch) != 0 {
			<-ch
		}
		time.Sleep(1 * time.Second)
		ch <- struct{}{}
	}
	//	<-timer2.C
	//	fmt.Println("Timer 2 stopped")

	//	time.Sleep(100 * time.Second)
}
