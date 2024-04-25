package main

import (
	"fmt"
)

const nnum = 10

func retNil(num *int) *int {
	if *num > nnum {
		return nil
	}

	return num
}

func main() {
	num := 11
	v := retNil(&num)

	fmt.Println(*v)
}
