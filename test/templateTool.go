package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	srcDirName := "./201708"
	targetDirName := "./temp"
	os.Mkdir(targetDirName, 0777)

	files := make([]string, 0)
	dirs, err := ioutil.ReadDir(srcDirName)

	if err != nil {
		return
	}

	for _, fi := range dirs {
		if fi.IsDir() {
			continue
		}

		fullName := srcDirName + string(os.PathSeparator) + fi.Name()
		targetName := targetDirName + string(os.PathSeparator) + fi.Name()
		fmt.Println(fullName)

		files = append(files, fullName)

		file, _ := os.Open(fullName)
		defer file.Close()

		data, _ := ioutil.ReadFile(file.Name())
		content := string(data)
		fmt.Println(content)

		// 替换一般的对象
		reg1 := regexp.MustCompile("\\$\\{[^\\{]*\\!?}")
		subStrs1 := reg1.FindAllString(content, -1)

		for _, subStr1 := range subStrs1 {
			fmt.Println(subStr1)

			repl := ""
			for i := 0; i < len(subStr1); i++ {
				repl += "&nbsp;"
			}
			if strings.Contains(subStr1, "good") || strings.Contains(subStr1, "basicNumberDictValue") {
				continue
			}
			content = reg1.ReplaceAllString(content, repl)
		}

		// 去掉签章
		reg3 := regexp.MustCompile("http://.*\\.png")
		subStrs3 := reg3.FindAllString(content, -1)

		for _, subStr3 := range subStrs3 {
			fmt.Println(subStr3)

			repl := ""

			content = reg3.ReplaceAllString(content, repl)
		}

		html, _ := os.Create(targetName)

		defer html.Close()

		html.WriteString(content)
	}
}
