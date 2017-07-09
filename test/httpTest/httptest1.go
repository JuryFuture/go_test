package main
import (
	"io"
	"os"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	defer resp.Body.Close()
	if err != nil {
		io.Copy(os.Stdout, resp.Body)
	}
}
