package core

import (
	"fmt"
	"log"
	"os"
)

func logging() {
	file, err := os.OpenFile("./app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	// log.SetOutput sets the oputput destination for the standard logger
	log.SetOutput(file)
	log.Println("Hello world")
}
