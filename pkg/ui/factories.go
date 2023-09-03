package ui

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

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
	list.TextStyle = termui.NewStyle(termui.ColorBlue)
	list.SelectedRowStyle = termui.NewStyle(termui.ColorGreen)
	list.WrapText = false
	return list
}

func createPopup() *widgets.Paragraph {
	popup := widgets.NewParagraph()
	termWidth, termHeight := termui.TerminalDimensions()
	popupWidth := termWidth / 3
	popupHeight := termHeight / 5
	popup.Title = "Popup"
	popup.TextStyle = termui.NewStyle(termui.ColorYellow)
	popup.WrapText = false
	popup.SetRect((termWidth-popupWidth)/2, (termHeight-popupHeight)/2, (termWidth+popupWidth)/2, (termHeight+popupHeight)/2)
	return popup
}
