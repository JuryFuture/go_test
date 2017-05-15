package main
import "fmt"
var (
	x int
	y int
)
var a int = 1
var b = 2
var c,d int = 3,4
var e,f = 5,6

func main() {
	x,y = 7,8
	a,h := 9,10
	const (
		i = 10
		j
	)	
	fmt.Println(a,b,c,d,e,f,h,x,y,i,j)
}
