package core

import (
	"fmt"
	"strings"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func input_output() {
	var s Shape = Rectangle{4.0, 6.0}
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())

}

func reader() {
	r := strings.NewReader("Learning is fun")
	buf := make([]byte, 4)
	for {
		n, err := r.Read(buf)
		fmt.Println(buf[:n], err)
		if err != nil {
			fmt.Println("Breaking out")
			break
		}
	}
}
