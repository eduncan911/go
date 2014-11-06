package cmd

import (
	"bytes"
	"os/exec"
)

func init() {
	osshell = &unixshell{}
}

type unixshell struct {
}

func (s *unixshell) Exec(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	var buf bytes.Buffer
	cmd.Stderr = &buf
	cmd.Stdout = &buf

	// assign return vars
	err := cmd.Run()
	out := buf.String()
	return out, err
}
