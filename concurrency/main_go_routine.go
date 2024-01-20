package main

import (
	"fmt"
	"time"
)

func main_goroutine() {
	go start()
	time.Sleep(1 * time.Second)
}

func start() {
	fmt.Println("In start")
}

func process() {
	fmt.Println("In process")
}
