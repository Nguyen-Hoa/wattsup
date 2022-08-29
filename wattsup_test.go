package wattsup

import (
	"bufio"
	"log"
	"os"
)

func main() {
	port := "ttyUSB0"
	command := []string{"./wattsup", port, "-g", "watts"}
	filename := "out.watts"
	args := WattsupArgs{port, filename, command}
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
