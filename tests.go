package main

import (
	"fmt"
)

func ret_nil(num *int) *int {
	if *num > 10 {
		return nil
	}
	return num
}

func main() {
	var num int = 2
	var v *int
	v = ret_nil(&num)
	fmt.Println(*v)
}
