package main
import "fmt"
import "reflect"
func main() {
	var arr1 [5]int
	arr2 := [5]int{1,2,3,4,5}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for _,v := range arr2 {
		fmt.Println(v)
	}

	arr3 := arr2
	arr3[1]++

	arr4 := &arr3
	fmt.Println(reflect.TypeOf(arr4))
	fmt.Println(arr2, arr3, *arr4)
}
