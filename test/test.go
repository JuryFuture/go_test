package main
import "fmt"
type A struct{
	a int
}

func main() {
	a := &A{1}
	b := a
	a = &A{2}
	fmt.Println(b)
}
