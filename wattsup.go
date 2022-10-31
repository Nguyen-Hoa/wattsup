package wattsup

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Wattsup struct {
	cmd      string
	path     string
	fullpath string
	running  bool

	cmd_  *exec.Cmd
	file_ *os.File
}

type WattsupArgs struct {
	Path string `json:"path"`
	Cmd  string `json:"cmd"`
	Name string `json:"name"`
}

func New(args WattsupArgs) *Wattsup {
	path := args.Path
	name := ""
	if args.Name != "" {
		name = args.Name
	} else {
		name = time.Now().Format("2006_01_02-15:04:05")
	}
	fullPath := fmt.Sprintf("./%s/%s.watts", path, name)
	file_, err := os.Create(fullPath)
	if err != nil {
		log.Print(err)
		return nil
	}

	proc, proc_args := strings.Split(args.Cmd, " ")[0], strings.Split(args.Cmd, " ")[1:]
	cmd_ := exec.Command(proc, proc_args...)
	cmd_.Stdout = file_

	w := Wattsup{}
	w.cmd = args.Cmd
	w.path = args.Path
	w.fullpath = fullPath
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
