package main
import "fmt"
func main() {
	var arr1 [5]int
	arr2 := []int{1,2,3,4,5}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for _,v := range arr2 {
		fmt.Println(v)
	}
}
