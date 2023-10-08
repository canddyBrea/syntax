package main

import (
	"fmt"

	"syntax/pkg"
)

func main() {
	// first tasks ***
	data := make([]int, 20, 40)
	for i := 0; i < 20; i++ {
		data[i] = i + 1
	}

	fmt.Println(cap(data), len(data))
	s1lice := pkg.Delete[int](data, 9, 10, 11, 12, 13, 14, 15, 16)
	fmt.Println(cap(s1lice), len(s1lice))
}
