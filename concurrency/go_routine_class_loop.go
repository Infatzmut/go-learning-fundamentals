package main

import (
	"fmt"
	"sync"
)

func go_routine_loop() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	fmt.Println("Done")
	wg.Wait()
}
