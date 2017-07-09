package main

import "fmt"

const DEFAULT_CAPACITY int = 10

/*
线性表
*/
type ArrayList struct {
	data [DEFAULT_CAPACITY]interface{}
	size int
}

//初始化
func NewArrayList() *ArrayList {
	li := &List{[DEFAULT_CAPACITY]interface{}{},0}
	return li
}

//增
func (li *ArrayList) Add(index int, o interface{}) {
	if li.size + 1 < DEFAULT_CAPACITY && index > -1 {
		if li.Get(index) != nil {
			for i := DEFAULT_CAPACITY - 1; i > index; i-- {
				li.data[i] = li.data[i-1]
			}
		}
		li.data[index] = o
		li.size ++
	} else {
		panic("out of index")
	}
}

//删
func (li *List) Delete(index int) interface{} {
	o := li.Get(index)

	for i:= index; i < DEFAULT_CAPACITY -1; i++ {
		li.data[i] = li.data[i+1]
	}
	li.size--
	li.data[li.size] = nil

	return o
}

//查
func (li *List) Get(index int) interface{} {
	if index < DEFAULT_CAPACITY && index > -1 {
		return li.data[index]
	} else {
		panic("out of index")
	}
}

//改
func (li *List) Set(index int, o interface{}) {
	if index < DEFAULT_CAPACITY && index > -1 {
		if li.Get(index) == nil {
			li.size ++
		}
		li.data[index] = o
	} else {
		panic("out of index")
	}
}

//计算长度
func (li *List) Size() int {
	return li.size
}

func main() {
	li := NewArrayList()
	li.Add(1,1)
	//fmt.Println(li.Get(0))
	defer fmt.Println(li)
	/*for i:=0; i < 10; i++ {
		li.Add(i)
	}*/
	//li.Delete(2)
	li.Set(7,2)
	li.Add(7,3)
	fmt.Println(li.Size())
}
