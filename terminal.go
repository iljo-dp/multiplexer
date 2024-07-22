// terminal.go
package main

import (
	"os/exec"

	"github.com/creack/pty"
)

type Terminal struct {
	ptmx *pty.FD // Change to *pty.FD
	cmd  *exec.Cmd
}

func NewTerminal() (*Terminal, error) {
	cmd := exec.Command("bash") // You can replace "bash" with any shell or command you prefer
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return nil, err
	}

	return &Terminal{
		ptmx: ptmx,
		cmd:  cmd,
	}, nil
}

func (t *Terminal) Resize(width, height int) {
	t.ptmx.Resize(width, height)
}

func (t *Terminal) Write(p []byte) (n int, err error) {
	return t.ptmx.Write(p)
}

func (t *Terminal) Read(p []byte) (n int, err error) {
	return t.ptmx.Read(p)
}

func (t *Terminal) Close() error {
	return t.ptmx.Close()
}
