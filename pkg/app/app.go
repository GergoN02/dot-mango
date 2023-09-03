package app

import (
	"sync"
	"time"

	"github.com/gizak/termui/v3"
	"github.com/thegenem0/dot-mango/pkg/events"
	"github.com/thegenem0/dot-mango/pkg/state"
	"github.com/thegenem0/dot-mango/pkg/ui"
)

func LaunchUI() {
	if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()

	state.GetAppState().LoadConfig()

	uiView := ui.NewView()
	uiView.SetLayout()
	uiView.Render()

	go events.HandleUserEvents(uiView)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		runEventLoop(uiView)
		wg.Done()
	}()

	wg.Wait()
}

func runEventLoop(uiView *ui.View) {
	var tick = time.NewTicker(10 * time.Millisecond)
	for {
		select {
		case <-tick.C:
			uiView.SetInfoBarText()
			uiView.SetConfigSelectorItems()
			uiView.Render()
		}
	}
}
