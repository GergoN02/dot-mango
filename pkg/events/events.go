package events

import (
	"fmt"
	"os"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/thegenem0/dot-mango/pkg/fileops"
	"github.com/thegenem0/dot-mango/pkg/models"
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
	KeyS
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
				if view.GetPopupType() == models.FileOverwritePopup {
					err := fileops.SymlinkSelectedDotfiles(view.GetSelectedMangoConfig().Path, view.GetActiveDotfileDirChildren())
					view.SetPopupType(models.InfoPopup)
					if err != nil {
						view.SetPopupContent(fmt.Sprintf("Error: %s", err))
					} else {
						view.SetPopupContent("Success!\n\nPress <Enter> to continue.")
					}

				}
			case "n", "N":
				view.TogglePopup()
			case "<Enter>":
				if view.GetPopupType() == models.InfoPopup {
					view.TogglePopup()
				}
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
			case "s":
				if err := fileops.CheckSymlinkPaths(view.GetActiveDotfileDirChildren()); err != nil {
					view.SetPopupContent("Some files already exist at the symlink locations.\nAre you sure you want to overwrite them?\n\nPress <y> to confirm, <n> to cancel.")
				} else {
					view.SetPopupContent("Are you sure you want to symlink these files? (y/n)\n\n")
				}
				view.SetPopupType(models.FileOverwritePopup)
				view.TogglePopup()
			}
		}
	}
}
