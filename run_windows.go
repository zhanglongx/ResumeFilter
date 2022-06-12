//go:build windows

package main

import (
	"os/exec"
)

func OpenFile(filename string) error {
	cmd := exec.Command("start", filename)
	return cmd.Run()
}
