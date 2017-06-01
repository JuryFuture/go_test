package main

import (
	"fmt"
	"sync"
	"strconv"
)

var total = 1000

/*
在goroutine内部调用wg.Add(1),可能会不起作用，可能还没执行之前主协程已经结束
func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)
	ch <- 0

	for i := 0; i < total; i++ {
		go func(name string, ch chan int) {
			wg.Add(1)
			defer wg.Done()
			a := <- ch
			a++
			fmt.Println(name, "计算后:", a)
			ch <- a
		}("goroutine" + strconv.Itoa(i), ch)
	}

	wg.Wait()

	fmt.Println("计算结果：", <-ch)
}*/
/*可以在goroutine之外调用wg.Add(),可以循环调用，也可以一次添加total个：:wq.Add(total)*/
func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)
	ch <- 0

	for i := 0; i < total; i++ {
		wg.Add(1)
		go func(name string, ch chan int) {
			a := <- ch
			a++
			fmt.Println(name, "计算后:", a)
			ch <- a
			defer wg.Done()
		}("goroutine" + strconv.Itoa(i), ch)
	}

	wg.Wait()

	fmt.Println("计算结果:", <-ch)

}
