package main

import (
	"log"
	"sync"
	"time"

	"github.com/gergon02/dot-mango/internal/events"
	"github.com/gergon02/dot-mango/internal/ui"
	"github.com/gergon02/dot-mango/internal/utils"
	"github.com/gizak/termui/v3"
)

var (
	uiEventChan chan events.UiEvent
)

func main() {
	if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()

	uiView := ui.NewView()
	uiView.SetLayout()
	uiView.Render()

	uiEventChan = make(chan events.UiEvent)

	config, err := utils.LoadConfig()
	if err != nil {
		log.Panicln("Your toml config is invalid", err)
	}

	go events.HandleUserEvents(uiEventChan)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		runEventLoop(uiView, config)
		wg.Done()
	}()

	wg.Wait()
}

func runEventLoop(uiView *ui.View, config *utils.AppConfig) {
	var tick = time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case e := <-uiEventChan:
			switch e {
			case events.KeyEsc:
				termui.Close()
				return
			default:
				println("default")
			}
		case <-tick.C:
			uiView.SetInfoBarText(config.UserConfig[0].Name + " " + config.UserConfig[0].Path)
		}
	}
}
