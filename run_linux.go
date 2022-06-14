//go:build linux

package main

import (
	"os/exec"
)

func OpenFile(filename string) error {
	return Run("xdg-open", filename)
}

func Pdf2txt(txt string, pdf string) error {
	return Run("pdf2txt", "-o", txt, pdf)
}

func Run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	return cmd.Run()
}
