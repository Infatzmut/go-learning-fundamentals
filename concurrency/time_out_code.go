package main

import (
	"fmt"
	"time"
)

func main2() {
	ch1 := make(chan int)
	go sendValue(ch1)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Select timeout")
	}
}

func sendValue(ch chan int) {
	// time.Sleep(3*time.Second)
	ch <- 10
}
