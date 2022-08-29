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

type WattsupArgs struct {
	port string
	file string
	cmd  []string
}

func (w *Wattsup) Init(args WattsupArgs) error {
	file_, err := os.Create(args.file)
	if err != nil {
		return errors.New("Failed to open file for watts output")
	}

	cmd_ := exec.Command(args.cmd[0], args.cmd[1:]...)
	cmd_.Stdout = file_

	w.port = args.port
	w.cmd = args.cmd
	w.file = args.file
	w.running = true

	w.cmd_ = cmd_
	w.file_ = file_

	return nil
}

func New(args WattsupArgs) *Wattsup {
	file_, err := os.Create(args.file)
	if err != nil {
		return nil
	}

	cmd_ := exec.Command(args.cmd[0], args.cmd[1:]...)
	cmd_.Stdout = file_

	w := Wattsup{}
	w.port = args.port
	w.cmd = args.cmd
	w.file = args.file
	w.running = true
	w.cmd_ = cmd_
	w.file_ = file_
	return &w
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
