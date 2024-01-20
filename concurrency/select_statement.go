package main

import (
	"fmt"
	"time"
)

func select_statement() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
		break
	case val2 := <-ch2:
		fmt.Println(val2)
	default:
		fmt.Println("Executed default block")
	}
	time.Sleep(1 * time.Second)
}

func goOne(ch chan string) {
	ch <- "Channel-2"
}

func goTwo(ch chan string) {
	ch <- "Channel-1"
}
