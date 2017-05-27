package main

import(
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"

	"mlib/library"
	"mp"
)

var lib = new(library.MusicManager)
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
		case "list":
			for i := 0; i < lib.Len(); i++ {
				m,_ := lib.Get(i)
				fmt.Println(i+1,":", m.Name, m.Artist, m.Source, m.Type)
			}
		case "add":
			if len(tokens) == 6 {
				id++
				lib.Add(&library.MusicEntry{strconv.Itoa(id),tokens[2],tokens[3],tokens[4],tokens[5]})
			} else {
				fmt.Println("USAGE: lib add <name><srtist><source><type>")
			}
		case "remove":
			if len(tokens) ==3 {
				index,_ := strconv.Atoi(tokens[2])
				lib.Remove(index)
			} else {
				fmt.Println("USAGE: lib remove <index>")
			}
		default:
			fmt.Println("Unrecognized lib command:",tokens[0])
	}
}

func handlePlayCommands(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e := lib.Find(tokens[1])


	if e ==nil {
		fmt.Println("The music", tokens[1], "not exist.")
	}

	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println("Enter following commands to control the player:\nlib list -- View the existing music lib\nlib add <name><artist><source><type> -- Add a music to the music lib\nlib remove <name> -- Remove the specified music from the lib\nplay <name> -- Play the specified music")

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter command-> ")

		rawLine,_,_ := r.ReadLine()
		line := string(rawLine)

		if line =="q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommands(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}
