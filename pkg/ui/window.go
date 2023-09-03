package ui

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/thegenem0/dot-mango/pkg/state"
)

func NewView() *View {
	var view = View{
		UiState: state.GetUiState(),
	}

	view.InfoBar = widgets.NewParagraph()
	view.ConfigSelector = createConfigSelector()
	view.DotfileDirs = createAvailableConfigs()

	view.Popup = createPopup()

	view.SetPanes([]termui.Drawable{
		view.InfoBar,
		view.ConfigSelector,
		view.DotfileDirs,
	})

	view.SetWidgetNames(map[string]termui.Drawable{
		"InfoBar":        view.InfoBar,
		"ConfigSelector": view.ConfigSelector,
		"DotfileDirs":    view.DotfileDirs,
		"Popup":          view.Popup,
	})

	setFocus(view.GetPanes()[view.GetHighlightedPane()])

	return &view
}

func (self *View) SetLayout() {
	self.Grid = termui.NewGrid()
	self.SetSize()
	self.Grid.Set(
		termui.NewCol(0.2,
			termui.NewRow(0.2, self.InfoBar),
			termui.NewRow(0.8, self.ConfigSelector),
		),
		termui.NewCol(0.8,
			termui.NewRow(1.0, self.DotfileDirs),
		),
	)
}

func (self *View) Render() {
	if self.GetPopupActive() {
		termui.Render(self.Grid, self.Popup)
	} else {
		termui.Render(self.Grid)
	}
}

func (self *View) SetSize() {
	termWidth, termHeight := termui.TerminalDimensions()
	self.Grid.SetRect(0, 0, termWidth, termHeight)
}

func (self *View) GetActiveWidgetName() string {
	for name, widget := range self.WidgetNames {
		if widget == self.ActivePane {
			return name
		}
	}
	return ""
}

func (self *View) SwitchPane() {
	removeFocus(self.GetPanes()[self.GetHighlightedPane()])
	self.SetHighlightedPane((self.GetHighlightedPane() + 1) % len(self.GetPanes()))
	setFocus(self.GetPanes()[self.GetHighlightedPane()])
	self.setActivePane()
}

func (self *View) TogglePopup() {
	self.SetPopupActive(!self.GetPopupActive())
}
