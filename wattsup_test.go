package wattsup

import (
	"bufio"
	"log"
	"os"
)

func main() {
	filename := "out.watts"
	command := "./wattsup ttyUSB0 -g watts"
	args := WattsupArgs{filename, command}
	m := Wattsup{}
	if err := m.Init(args); err != nil {
		m.Start()
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
