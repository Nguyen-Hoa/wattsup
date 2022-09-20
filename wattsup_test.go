package wattsup

import (
	"bufio"
	"log"
	"os"
)

func main() {
	path := "./"
	command := "./wattsup ttyUSB0 -g watts"
	args := WattsupArgs{Path: path, Cmd: command}
	m := New(args)
	if m != nil {
		m.Start()
	} else {
		log.Panic("Failed to start meter")
	}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "quit" {
			break
		}
		log.Println(input.Text())
	}

	m.Stop()
}
