package mp

import(
	"fmt"
	"time"
)

type MP3Player struct {
}

func (p *MP3Player) Play(source string) {
	fmt.Println("Playing MP3 music",source)
	progress := 0
	for progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		progress += 10
	}

	fmt.Println("\nFinished playing", source)
}
