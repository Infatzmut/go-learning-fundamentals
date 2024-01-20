package core

import (
	"errors"
	"fmt"
)

func process(i int) error {
	if i%2 == 0 {
		//return errors.New("ONly odd numbers allowed")
		return fmt.Errorf("Only odd numbers allowed, got: %d", i)
	}
	return nil
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("Operation successful")
}
func error_method() {
	// New erturns and error that formats as the given text
	err := errors.New("Custom error ocurred")
	fmt.Println(err)

	err = process(3)
	checkError(err)
	err = process(2)
	checkError(err)
}
