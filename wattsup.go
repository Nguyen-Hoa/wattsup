package wattsup

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Wattsup struct {
	cmd     string
	file    string
	running bool

	cmd_  *exec.Cmd
	file_ *os.File
}

type WattsupArgs struct {
	File string
	Cmd  string
}

func New(args WattsupArgs) *Wattsup {
	file_, err := os.Create(args.File)
	if err != nil {
		return nil
	}

	proc, proc_args := strings.Split(args.Cmd, " ")[0], strings.Split(args.Cmd, " ")[1:]
	cmd_ := exec.Command(proc, proc_args...)
	cmd_.Stdout = file_

	w := Wattsup{}
	w.cmd = args.Cmd
	w.file = args.File
	w.running = false
	w.cmd_ = cmd_
	w.file_ = file_
	return &w
}

func (w *Wattsup) Start() error {
	if err := w.cmd_.Start(); err != nil {
		log.Fatal(err)
		return errors.New("Failed to start power meter")
	}
	w.running = true
	log.Println("Started meter...")
	return nil
}

func (w *Wattsup) Stop() error {
	if err := w.cmd_.Process.Kill(); err != nil {
		log.Fatal(err)
		return errors.New("Failed to stop meter")
	}
	w.file_.Close()
	w.running = false
	return nil
}

func (w *Wattsup) Running() bool {
	if w.running {
		return true
	} else {
		return false
	}
}
