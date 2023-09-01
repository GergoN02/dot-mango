package main

import (
	"sync"
	"time"

	"github.com/gizak/termui/v3"
	"github.com/thegenem0/dot-mango/internal/events"
	"github.com/thegenem0/dot-mango/internal/ui"
	"github.com/thegenem0/dot-mango/internal/utils"
)

func main() {
	if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()

	appConfig, err := utils.LoadConfig()
	if err != nil {
		panic(err)
	}

	uiView := ui.NewView(appConfig)
	uiView.SetLayout()
	uiView.Render()

	go events.HandleUserEvents(uiView, appConfig)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		runEventLoop(uiView, appConfig)
		wg.Done()
	}()

	wg.Wait()
}

func runEventLoop(uiView *ui.View, config *utils.AppConfig) {
	var tick = time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-tick.C:
			uiView.SetInfoBarText(config.UserConfig[0].Name + " " + config.UserConfig[0].Path)
			uiView.SetConfigSelectorItems(utils.GetConfigNames(config))
			uiView.SetAvailableConfigsItems(config.UserConfig[0].Path)
			uiView.Render()
		}
	}
}
