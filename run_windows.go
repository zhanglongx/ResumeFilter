//go:build windows

package main

import (
	"os/exec"
)

func OpenFile(filename string) error {
	return Run("explorer", filename)
}

func Pdf2txt(txt string, pdf string) error {
	return Run("pdf2txt.exe", "-o", txt, pdf)
}

func Run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Run()
	return nil
}
