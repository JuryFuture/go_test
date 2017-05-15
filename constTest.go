package main
import "fmt"
func main() {
	const (
		a = iota
		b
		c
	)
	const (
		d = 10
		e
		f
	)

	fmt.Println(a,b,c,d,e,f)

}
