package main

import (
	"fmt"
	"sync"
	"time"
)

func calculateSquares(i int, wg *sync.WaitGroup) {
	fmt.Println(i * i)
	wg.Done()
}

func wait_groups() {
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go calculateSquares(i, &wg)
	}
	elapsed := time.Since(start)
	wg.Wait()
	fmt.Println("Function took ", elapsed)
}
