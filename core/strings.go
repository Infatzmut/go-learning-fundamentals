package core

import (
	"fmt"
	"strings"
)

func main() {
	learning := "learning standard library in Go"
	fun := "library in Go"
	result := strings.Contains(learning, fun)
	fmt.Println(result)
	result_count := strings.Count(learning, fun)
	fmt.Println(result_count)
}
