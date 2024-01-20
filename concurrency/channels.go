package main

import (
	"fmt"
	"time"
)

func channels() {

	ch := make(chan string)
	go sell(ch)
	go buy(ch)
	time.Sleep(2 * time.Second)
}

// sends data to the channel
// is send as a reference implicitly
func sell(ch chan string) {
	ch <- "Furniture"
	fmt.Println("Sent data to the channel")
}

// receive data from the channel
func buy(ch chan string) {
	fmt.Println("Witing for data")
	val := <-ch
	fmt.Println("Received data - ", val)
}
