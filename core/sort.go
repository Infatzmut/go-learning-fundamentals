package core

import (
	"fmt"
	"sort"
)

func sorting() {
	vars := []int{5, 0, 2, 3, 4, 9, 6}
	sort.Ints(vars)
	fmt.Println(vars)

	vars2 := []string{"Learning", "Golang", "on", "Kodekloud"}
	sort.Strings(vars2)
	fmt.Println(vars2)
}
