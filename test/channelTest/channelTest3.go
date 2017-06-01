package main

import "fmt"

import "strconv"

func count(name string, ch, done chan int) {
	a := <-ch
	a++
	ch <- a

	fmt.Println(name,"计算后：",a)

	done <- 1
}

func main() {
	total := 1000
	ch := make(chan int, 1)
	ch <- 0

	dones := make([]chan int, total)
	for i:= 0; i < total; i++ {
		done := make(chan int, 1)
		dones[i] = done
		go count("goroutine" + strconv.Itoa(i), ch, done)
	}

	for j := 0; j < total; j++ {
		<-dones[j]
	}

	fmt.Println("结果值：", <-ch)
}
