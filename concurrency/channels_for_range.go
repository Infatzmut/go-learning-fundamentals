package main

import (
	"fmt"
	"sync"
)

func range2() {
	// for range
	nums := [5]int{1, 2, 3, 4, 5}
	for i, item := range nums {
		fmt.Println(i, item)
	}
}

func principal() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go sell2(ch, &wg)
	go buy2(ch, &wg)
	wg.Wait()
}

func buy2(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for data")
	for val := range ch {
		fmt.Println("Recieved", val)
	}
	wg.Done()
}

func sell2(ch chan int, wg *sync.WaitGroup) {
	ch <- 1
	ch <- 2
	fmt.Println("Sent all data")
	close(ch)
	wg.Done()
}
