package events

import (
	"github.com/gizak/termui/v3"
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
	KeyQ
	KeyH
	KeyJ
	KeyK
	KeyL
)

func HandleUserEvents(uiEventChan chan UiEvent) {
	uiEvents := termui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>", "<Escape>":
				println("Exiting...")
				uiEventChan <- KeyEsc
			case "<Resize>":
				println("Resizing...")
				uiEventChan <- KeySpace
			}
		}
	}
}
