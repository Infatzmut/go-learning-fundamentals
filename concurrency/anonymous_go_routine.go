package main

import (
	"fmt"
	"time"
)

// anonymous funcitons are those functions that dont have any name
// Simply put, anonymous functions don't use any varialbes as a name when they are declared, and can also be called using go-routine

func anonymous_go_routine() {
	go func() {
		fmt.Println("In anonymous method")
	}()
	time.Sleep(1 * time.Second)
}
