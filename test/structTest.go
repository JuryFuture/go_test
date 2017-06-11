package main
import "fmt"
type User struct {
	Name string
	Age int
}

func (user *User) SayHello() {
	fmt.Printf("Hello,%s!\n", user.Name)
}

func main() {
	var u = new(User)
	u.Name = "Jack"
	u.SayHello()

	u1 := &User{}
	u1.Name = "Tom"
	u1.SayHello()

	u2 := &User{"Jury", 10}
	u2.SayHello()

	u3 := &User{Name:"Future"}
	u3.SayHello()


	u4 := u
	u4.Name = "Jury"
	u.SayHello()
	u4.SayHello()
}
