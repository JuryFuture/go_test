package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
			case ch <- 1:
				fmt.Println("Send 1")
			case ch <- 0:
				fmt.Println("Send 0")
		}

		a := <-ch
		fmt.Println(a)
	}
}
