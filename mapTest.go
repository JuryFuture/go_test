package main
import (
	"fmt"
)
func main() {
	var m1 map[string] string = make(map[string] string)
	m1["1"] = "1"
	fmt.Println(m1["1"])

	m2 := map[string] string{"2":"2","3":"3"}
	fmt.Println(m2["2"],m2["3"])
}
