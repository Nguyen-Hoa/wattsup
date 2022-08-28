package main

import (
	"bufio"
	"log"
	"os"
	"watts-up-go/meter"
)

func main() {
	port := "ttyUSB0"
	command := []string{"./wattsup", port, "-g", "watts"}
	filename := "out.watts"
	m := meter.New(port, filename, command)
	m.Start()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "quit" {
			break
		}
		log.Println(input.Text())
	}

	m.Stop()
}
