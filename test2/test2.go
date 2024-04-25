package main

import (
	"fmt"
)

func main() {
	var p *int
	var num int = *p
	fmt.Println(num)
}
