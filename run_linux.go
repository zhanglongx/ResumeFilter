//go:build linux

package main

import (
	"os/exec"
)

func OpenFile(filename string) error {
	cmd := exec.Command("xdg-open", filename)
	return cmd.Run()
}
