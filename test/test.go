package main
import "fmt"
type A struct{
	a int
}

func main() {
	a := &A{1}
	fmt.Println(a.a)
}
