package ui

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/thegenem0/dot-mango/pkg/state"
)

func (self *View) countSelectedConfigItems() int {
	var count int
	for _, config := range state.GetUiState().GetActiveDotfileDirChildren() {
		if config.Selected {
			count++
		}
	}
	return count
}

func (self *View) setActivePane() {
	uiState := state.GetUiState()

	uiState.SetActivePane(uiState.GetPanes()[uiState.GetHighlightedPane()])
}

func (self *View) removeActivePane() {
	uiState := state.GetUiState()

	uiState.SetActivePane(nil)
}

func setFocus(w termui.Drawable) {
	switch widget := w.(type) {
	case *widgets.List:
		widget.BorderStyle.Fg = termui.ColorRed
	case *widgets.Paragraph:
		widget.BorderStyle.Fg = termui.ColorRed
	}
}

func removeFocus(w termui.Drawable) {
	switch widget := w.(type) {
	case *widgets.List:
		widget.BorderStyle.Fg = termui.ColorWhite
	case *widgets.Paragraph:
		widget.BorderStyle.Fg = termui.ColorWhite
	}
}
