package main
import "fmt"
func main(){
	var str string
	str = "Hello,World! 你好世界"
	for i:=0;i<len(str);i++ {
		fmt.Printf("%s的地%d个字符为%c\n",str,i,str[i])
	}
	for j,s := range str {
		fmt.Println(j,s)
	}
}
