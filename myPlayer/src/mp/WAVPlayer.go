package mp

import(
	"fmt"
	"time"
)

type WAVPlayer struct {
}

func (p * WAVPlayer) Play(source string) {
	fmt.Println("Playing WAV music", source)

	progress := 0

	for progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		progress += 10
	}

	fmt.Println("\nFinished playing", source)
}
