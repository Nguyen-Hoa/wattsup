package wattsup

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

type Wattsup struct {
	port    string
	cmd     []string
	file    string
	running bool

	cmd_  *exec.Cmd
	file_ *os.File
}

func (w *Wattsup) Init(port, file string, cmd []string) error {
	file_, err := os.Create(file)
	if err != nil {
		return errors.New("Failed to open file for watts output")
	}

	cmd_ := exec.Command(cmd[0], cmd[1:]...)
	cmd_.Stdout = file_

	w.port = port
	w.cmd = cmd
	w.file = file
	w.running = true

	w.cmd_ = cmd_
	w.file_ = file_

	return nil
}

func (w *Wattsup) Start() error {
	if err := w.cmd_.Start(); err != nil {
		log.Fatal(err)
		return errors.New("Failed to start power meter")
	}
	log.Println("Started meter...")
	return nil
}

func (w *Wattsup) Stop() error {
	if err := w.cmd_.Process.Kill(); err != nil {
		log.Fatal(err)
		return errors.New("Failed to stop meter")
	}
	w.file_.Close()
	return nil
}

func (w *Wattsup) Running() bool {
	if w.running {
		return true
	} else {
		return false
	}
}
