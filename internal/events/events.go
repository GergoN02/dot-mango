package events

import (
	"os"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/thegenem0/dot-mango/internal/ui"
	"github.com/thegenem0/dot-mango/internal/utils"
)

type UiEvent int

const (
	KeyArrowUp UiEvent = 1 << iota
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
	KeyEsc
	KeyEnter
	KeySpace
	KeyTab
	KeyQuit
	KeyQ
	KeyH
	KeyJ
	KeyK
	KeyL
)

func HandleUserEvents(view *ui.View, config *utils.AppConfig) {
	uiEvents := termui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "<C-c>":
			termui.Close()
			os.Exit(0)
		case "<Tab>":
			view.SwitchPane()
		case "<Resize>":
			println("Resizing...")
			view.SetSize()
		}

		if view.ActivePane != nil {
			switch view.ActivePane.(type) {
			case *widgets.List:
				switch e.ID {
				case "j", "<Down>":
					view.ActivePane.(*widgets.List).ScrollDown()
				case "k", "<Up>":
					view.ActivePane.(*widgets.List).ScrollUp()
				case " ":

				}
			}
		}
	}
}
