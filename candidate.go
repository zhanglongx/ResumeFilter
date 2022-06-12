package main

import (
	"bytes"
	"os"
	"regexp"
	"strings"

	"github.com/dslipak/pdf"
)

type Candidate struct {
	Path    string
	College string
}

func (c *Candidate) Parse() error {
	f, err := os.Open(c.Path)
	if err != nil {
		return err
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return err
	}

	r, err := pdf.NewReader(f, fi.Size())
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return err
	}

	_, err = buf.ReadFrom(b)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`\p{Han}{2,4}(大学|学院)`)
	Colleges := re.FindAllString(buf.String(), -1)

	c.College = uniqCollege(Colleges)

	return nil
}

func uniqCollege(list []string) string {
	var t []string
	for _, s := range list {
		var bFound bool
		for _, ss := range t {
			if s == ss {
				bFound = true
			}
		}

		if !bFound {
			t = append(t, s)
		}
	}

	return strings.Join(t, "_")
}
