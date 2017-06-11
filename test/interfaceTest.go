package main
import (
	"fmt"
)

type ReadWriter interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}

type IOStream interface {
	Write(buf []byte) (n int, err error)
	Read(buf []byte) (n int, err error)
}

type Reader interface {
	Read(buf []byte) (n int, err error)
}

type File struct {
}

func (file *File) Read(buf []byte) (n int, err error) {
	fmt.Println("read")
	return 0,nil
}

func (file *File) Write(buf []byte) (n int, err error) {
	fmt.Println("write")
	return 0,nil
}

func main() {
	buf := make([]byte, 0)

	var f1 ReadWriter = new(File)
	fmt.Println(f1) //输出：&{}
	f1.Read(buf)
	f1.Write(buf)

	var f2 IOStream = &File{}
	fmt.Println(f2);
	f2.Read(buf)
	f2.Write(buf)

	var f3 Reader = f2
	f3.Read(buf)
}
