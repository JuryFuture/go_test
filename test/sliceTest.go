package main
import "fmt"
func main(){
	arr := []int{0,1,2,3,4,5,6,7,8,9}
	//通过现有数组创建
	var sl1 []int = arr[:5]
	//直接创建
	sl2 := make([]int,5)
	var sl3 []int = make([]int,5,10)
	//通过现有slice创建
	sl4 := sl3[:9]

	for i:=0; i<len(sl1);i++{
		fmt.Println(sl1[i])
	}

	for j:=0;j<len(sl2);j++{
		fmt.Println(sl2[j])
	}

	fmt.Println("cap()",cap(sl3))
	fmt.Println("len()",len(sl3))

	for _,v := range sl4 {
		fmt.Println(v)
	}
}
