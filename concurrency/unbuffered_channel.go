package main

import (
	"fmt"
	"sync"
)

func unbuffered() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 3)
	go sels(ch, &wg)
	wg.Wait()
}

func sels(ch chan int, s *sync.WaitGroup) {
	ch <- 10
	ch <- 11
	ch <- 12
	go buys(ch, s)
	fmt.Println("Sent all data to the channel")
	s.Done()
}

func buys(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Witing for data")
	fmt.Println("Received data", <-ch)
	wg.Done()
}
