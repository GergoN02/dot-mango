package ui

import (
	"log"
	"os"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func RunUi() {
	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	currentDir := "/home/gergon02/"

	list := widgets.NewList()
	list.Title = "List"
	list.TextStyle = termui.NewStyle(termui.ColorYellow)
	list.WrapText = false
	list.SetRect(0, 0, 25, 8)

	updateFileList(currentDir, list)

	termui.Render(list)

	uiEvents := termui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			list.ScrollDown()
		case "k", "<Up>":
			list.ScrollUp()
		}

		termui.Render(list)
	}
}

func updateFileList(dir string, list *widgets.List) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("failed to read directory: %v", err)
	}

	var rows []string
	for _, file := range files {
		if file.IsDir() {
			rows = append(rows, file.Name()+"/")
		} else {
			rows = append(rows, file.Name())
		}
	}

	list.Rows = rows
	if len(rows) > 0 {
		list.SelectedRow = 0
	}
}
