package message

import (
	"fmt"
	"testing"
)

func printSlice(s []int) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func TestSlice(t *testing.T) {
	s := make([]int, 2, 10)

	s = append(s, 1)
	printSlice(s)
	test(s)
	fmt.Println("=======")
	printSlice(s)
}

func test(s []int) {
	// s = append(s, 10)
	s[1] = 10
}
