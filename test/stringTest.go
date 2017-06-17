package main
import "fmt"
func main(){
	var str string
	str = "Hello,World! 你好世界"
	// byte类型
	for i:=0;i<len(str);i++ {
		fmt.Printf("%s的地%d个字符为%c\n",str,i,str[i])
	}
	
	// rune类型
	for j,s := range str {
		fmt.Println(j,s)
	}

	// 输出中文字符
	for t,n := range str {
		fmt.Println(t,string(n))
	}
}
