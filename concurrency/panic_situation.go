package main

import "fmt"

func panic_situation() {
	ch := make(chan int, 10)
	ch <- 10
	ch <- 11
	val, ok := <-ch
	fmt.Println(val, ok)
	close(ch)
	close(ch)
}
