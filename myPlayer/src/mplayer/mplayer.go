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

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tockens[1] {
		case "list":
			for i := 0; i < lib.Len(); i++ {
				m,_ := lib.Get(i)
				fmt.Println(i+1,":", m.Name, m.Artist, m.Source, m.Type)
			}
		case "add":
			if len(tockens) == 6 {
				id++
				lib.Add(&MusicEntry{id,tockens[2],tockens[3],tockens[4],tockens[5]})
			} else {
				fmt.Println("USAGE: lib add <name><srtist><source><type>")
			}
		case "remove":
			if len(tockens) ==3 {
				lib.Remove(strconv.Atoi(tockens[2]))
			} else {
				fmt.Println("USAGE: lib remove <index>")
			}
		default:
			fmt.Println("Unrecognized lib command:",tockens[0])
	}
}

func handlePlayCommand(tockens []string) {
	if len(tockens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e := lib.Find(tockens[1])


	if e ==nil {
		fmt.Println("The music", tockens[1], "not exist.")
	}

	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println("Enter following commands to control the player:\nlib list -- View the existing music lib\nlib add <name><artist><source><type> -- Add a music to the music lib\nlib remove <name> -- Remove the specified music from the lib\nplay <name> -- Play the specified music")

	lib := library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter command-> ")

		rawLine,_,_ := r.ReadLine()
		line := string(rawLine)

		if line =="q" || line == "e" {
			break
		}

		tockens := strings.Split(line, " ")

		if tockens[0] == "lib" {
			handleLibCommand(tockens)
		} else if tockens[0] == "play" {
			handlePlayCommand(tockens)
		} else {
			fmt.Println("Unrecognized command:", tockens[0])
		}
	}
}
