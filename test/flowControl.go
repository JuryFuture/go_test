package main

import (
	"fmt"
)

func main() {
	i := 1
	if i == 1 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	switch i {
	case 0:
		fmt.Println(0)
	case 2:
		fallthrough
	case 1:
		fmt.Println(1)
	default:
		fmt.Println("default")
	}

	for j := 1; j < 10; j++ {
		fmt.Println(j)
	}
}
