package events

import (
	"os"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/thegenem0/dot-mango/pkg/ui"
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

func HandleUserEvents(view *ui.View) {
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
		activePaneName := view.GetActiveWidgetName()

		switch activePaneName {
		case "Popup":
			switch e.ID {
			case "<Escape>", "q":
				view.TogglePopup()
			case "y", "Y":
				view.TogglePopup()
			default:
				break
			}
		case "ConfigSelector":
			switch e.ID {
			case "j", "<Down>":
				view.ActivePane.(*widgets.List).ScrollDown()
			case "k", "<Up>":
				view.ActivePane.(*widgets.List).ScrollUp()
			case "<Enter>":
				view.SetActiveMangoConfig()
			}
		case "DotfileDirs":
			switch e.ID {
			case "j", "<Down>":
				view.ActivePane.(*widgets.List).ScrollDown()
			case "k", "<Up>":
				view.ActivePane.(*widgets.List).ScrollUp()
			case "<Space>":
				view.ToggleSelectedConfigItem()
			case "a":
				view.ToggleAllConfigItems()
			case "p":
				view.TogglePopup()
			}
		}
	}
}
