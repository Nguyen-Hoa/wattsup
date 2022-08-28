package meter

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

type Meter struct {
	port    string
	cmd     []string
	file    string
	running bool

	cmd_  *exec.Cmd
	file_ *os.File
}

func New(port, file string, cmd []string) (*Meter, error) {
	m := Meter{}
	file_, err := os.Create(file)
	if err != nil {
		return &m, errors.New("Failed to open file for watts output")
	}

	cmd_ := exec.Command(cmd[0], cmd[1:]...)
	cmd_.Stdout = file_

	m = Meter{port, cmd, file, true, cmd_, file_}
	return &m, nil
}

func (m *Meter) Start() error {
	if err := m.cmd_.Start(); err != nil {
		log.Fatal(err)
		return errors.New("Failed to start power meter")
	}
	log.Println("Started meter...")
	return nil
}

func (m *Meter) Stop() error {
	if err := m.cmd_.Process.Kill(); err != nil {
		log.Fatal(err)
		return errors.New("Failed to stop meter")
	}
	m.file_.Close()
	return nil
}

func (m *Meter) Running() bool {
	if m.running {
		return true
	} else {
		return false
	}
}
