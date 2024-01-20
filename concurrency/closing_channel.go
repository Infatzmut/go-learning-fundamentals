package main

import "fmt"

func closing_channel() {
	ch := make(chan int, 10)
	ch <- 10
	ch <- 11
	val, ok := <-ch
	fmt.Println(val, ok)
	close(ch)
	val, ok = <-ch
	fmt.Println(val, ok)
	val, ok = <-ch
	fmt.Println(val, ok)
}
