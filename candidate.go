package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

type Candidate struct {
	Path    string
	College string
}

func (c *Candidate) Parse() error {
	txt := strings.TrimSuffix(c.Path, filepath.Ext(c.Path))

	err := Pdf2txt(txt, c.Path)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile(txt)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`\p{Han}{2,4}(大学|学院)`)
	Colleges := re.FindAllString(string(b), -1)

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
