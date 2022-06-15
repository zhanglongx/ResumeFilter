package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/mholt/archiver/v4"
	"github.com/zhanglongx/ResumeFilter/theme"
)

const (
	APP_NAME = "ResumeFilter"
	VERSION  = "1.1.2"
)

type CandController struct {
	IsCheck bool
	Cand    Candidate
}

type ArchiverController struct {
	CandControllers []*CandController

	fsys   fs.FS
	tmpDir string
}

func main() {
	optVer := flag.Bool("version", false, "print version")

	flag.Parse()

	optFiles := flag.Args()

	if *optVer {
		fmt.Printf("%s %s", APP_NAME, VERSION)
		os.Exit(1)
	}

	if len(optFiles) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	} else if len(optFiles) > 1 {
		log.Fatal()
	}

	a := &ArchiverController{}

	err := a.Open(optFiles[0])
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := a.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = a.Parse()
	if err != nil {
		log.Fatal(err)
	}

	// GUI
	app := app.New()
	app.Settings().SetTheme(&theme.MyTheme{})

	win := app.NewWindow(APP_NAME)

	// lets show GUI in parent window
	win.SetContent(a.makeUI())
	win.ShowAndRun()
}

func (c *CandController) OnCheck(isCheck bool) {
	c.IsCheck = isCheck
}

func (c *CandController) OnButton() {
	err := OpenFile(c.Cand.Path)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *ArchiverController) Open(filename string) error {
	fsys, err := archiver.FileSystem(filename)
	if err != nil {
		return err
	}

	a.fsys = fsys

	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}

	a.tmpDir = tmp

	return nil
}

func (a *ArchiverController) Close() error {
	return os.RemoveAll(a.tmpDir)
}

func (a *ArchiverController) Parse() error {
	var files []string
	err := fs.WalkDir(a.fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if strings.ToLower(filepath.Ext(path)) != ".pdf" {
			log.Printf("not supported file format: %s", path)
			return nil
		}

		// fmt.Println("Walking:", path, "Dir?", d.IsDir())

		src, err := a.fsys.Open(path)
		if err != nil {
			return err
		}

		defer src.Close()

		// FIXME: duplicated is skipped
		dst, err := os.Create(filepath.Join(a.tmpDir, filepath.Base(path)))
		if err != nil {
			return err
		}

		defer dst.Close()

		_, err = io.Copy(dst, src)
		if err != nil {
			return err
		}

		files = append(files, dst.Name())

		return nil
	})

	if err != nil {
		return err
	}

	// XXX: fs.WalkDir() get randomly
	sort.Strings(files)
	for _, f := range files {
		cc := &CandController{Cand: Candidate{Path: f}}
		err := cc.Cand.Parse()
		if err != nil {
			return err
		}

		a.CandControllers = append(a.CandControllers, cc)
	}

	return nil
}

func (a *ArchiverController) makeUI() *fyne.Container {
	var row []fyne.CanvasObject

	for _, c := range a.CandControllers {
		check := widget.NewCheck("", c.OnCheck)

		btn := widget.NewButton("打开", c.OnButton)

		name := widget.NewLabel(filepath.Base(c.Cand.Path))
		college := widget.NewLabel(c.Cand.College)

		row = append(row, container.NewHBox(
			check,
			btn,
			name,
			college,
		))
	}

	output := widget.NewMultiLineEntry()

	runBtn := widget.NewButton("输出", func() {
		var outText string
		for _, c := range a.CandControllers {
			if c.IsCheck {
				outText += fmt.Sprintf("%s\n", filepath.Base(c.Cand.Path))
			}
		}

		output.SetText(outText)
	})

	return container.NewVBox(
		container.NewGridWithColumns(1,
			row...,
		),
		runBtn,
		output)
}
