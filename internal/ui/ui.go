package ui

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type View struct {
	Grid    *termui.Grid
	InfoBar *widgets.Paragraph
}

func NewView() *View {
	var view = View{}

	view.InfoBar = widgets.NewParagraph()

	return &view
}

func (v *View) SetLayout() {
	v.Grid = termui.NewGrid()
	v.SetSize()
	v.Grid.Set(
		termui.NewRow(1.0, v.InfoBar),
	)
}

func (v *View) Render() {
	termui.Render(v.Grid)
}

func (v *View) SetSize() {
	termWidth, termHeight := termui.TerminalDimensions()
	v.Grid.SetRect(0, 0, termWidth, termHeight)
}

func (v *View) SetInfoBarText(text string) {
	v.InfoBar.Title = "Current Config"
	v.InfoBar.Text = text
	v.Render()
}
