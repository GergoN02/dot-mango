package ui

import (
	"fmt"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/thegenem0/dot-mango/internal/utils"
)

type View struct {
	Grid            *termui.Grid
	InfoBar         *widgets.Paragraph
	ConfigSelector  *widgets.List
	UserConfigDirs  *widgets.List
	highlightedPane int
	ActivePane      termui.Drawable
	panes           []termui.Drawable
}

func createConfigSelector() *widgets.List {
	list := widgets.NewList()
	list.Title = "Config List"
	list.TextStyle = termui.NewStyle(termui.ColorYellow)
	list.WrapText = false
	return list
}

func createAvailableConfigs() *widgets.List {
	list := widgets.NewList()
	list.Title = "Available Configs"
	list.TextStyle = termui.NewStyle(termui.ColorYellow)
	list.WrapText = false
	return list
}

func NewView(appConfig *utils.AppConfig) *View {
	var view = View{}

	view.InfoBar = widgets.NewParagraph()
	view.ConfigSelector = createConfigSelector()
	view.UserConfigDirs = createAvailableConfigs()

	view.panes = []termui.Drawable{
		view.InfoBar,
		view.ConfigSelector,
		view.UserConfigDirs,
	}
	view.highlightedPane = 0
	setFocus(view.panes[view.highlightedPane])

	return &view
}

func (v *View) SetLayout() {
	v.Grid = termui.NewGrid()
	v.SetSize()
	v.Grid.Set(
		termui.NewCol(0.2,
			termui.NewRow(0.2, v.InfoBar),
			termui.NewRow(0.8, v.ConfigSelector),
		),
		termui.NewCol(0.8,
			termui.NewRow(1.0, v.UserConfigDirs),
		),
	)
}

func (v *View) Render() {
	termui.Render(v.Grid)
}

func (v *View) SetSize() {
	termWidth, termHeight := termui.TerminalDimensions()
	v.Grid.SetRect(0, 0, termWidth, termHeight)
}

func (v *View) SwitchPane() {
	removeFocus(v.panes[v.highlightedPane])
	v.highlightedPane = (v.highlightedPane + 1) % len(v.panes)
	setFocus(v.panes[v.highlightedPane])
	v.setActivePane()
}

func (v *View) SetInfoBarText(text string) {
	v.InfoBar.Title = "Current Config"
	v.InfoBar.Text = text
}

func (v *View) SetConfigSelectorItems(items []string) {
	v.ConfigSelector.Rows = items
}

func (v *View) SetAvailableConfigsItems(path string) {
	configDirs, err := utils.GetUserConfigDirs(path)
	if err != nil {
		panic(err)
	}

	v.UserConfigDirs.Rows = make([]string, len(configDirs))
	for i, config := range configDirs {
		prefix := "[ ]"
		if config.Selected {
			prefix = "[x]"
		}
		v.UserConfigDirs.Rows[i] = fmt.Sprintf("%s %s -> %s", prefix, config.FolderName, config.SymlinkTarget)
	}
}

func (v *View) setActivePane() {
	v.ActivePane = v.panes[v.highlightedPane]
}

func (v *View) removeActivePane() {
	v.ActivePane = nil
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
