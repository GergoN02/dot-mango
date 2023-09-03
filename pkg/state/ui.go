package state

import (
	"sync"

	"github.com/gizak/termui/v3"
	"github.com/thegenem0/dot-mango/pkg/types"
)

var (
	uiInstance *UiState
	uiOnce     sync.Once
)

type UiState struct {
	ActiveDotfileDirChildren []types.DotfileDirectory

	// UI Management
	ActivePane  termui.Drawable
	panes       []termui.Drawable
	WidgetNames map[string]termui.Drawable

	// Current States or Values
	highlightedPane     int
	popupActive         bool
	selectedMangoConfig types.MangoConfig
}

func GetUiState() *UiState {
	uiOnce.Do(func() {
		uiInstance = &UiState{
			ActiveDotfileDirChildren: []types.DotfileDirectory{},
			panes:                    []termui.Drawable{},
			WidgetNames:              map[string]termui.Drawable{},
			highlightedPane:          0,
		}
	})
	return uiInstance
}

func (self *UiState) SetSelectedMangoConfig(mangoConfig types.MangoConfig) {
	self.selectedMangoConfig = mangoConfig
}

func (self *UiState) GetSelectedMangoConfig() types.MangoConfig {
	return self.selectedMangoConfig
}

func (self *UiState) SetActiveDotfileDirChildren(dotfileDirChildren []types.DotfileDirectory) {
	self.ActiveDotfileDirChildren = dotfileDirChildren
}

func (self *UiState) GetActiveDotfileDirChildren() []types.DotfileDirectory {
	return self.ActiveDotfileDirChildren
}

func (self *UiState) SetActivePane(pane termui.Drawable) {
	self.ActivePane = pane
}

func (self *UiState) GetActivePane() termui.Drawable {
	return self.ActivePane
}

func (self *UiState) SetPanes(panes []termui.Drawable) {
	self.panes = panes
}

func (self *UiState) GetPanes() []termui.Drawable {
	return self.panes
}

func (self *UiState) SetWidgetNames(widgetNames map[string]termui.Drawable) {
	self.WidgetNames = widgetNames
}

func (self *UiState) GetWidgetNames() map[string]termui.Drawable {
	return self.WidgetNames
}

func (self *UiState) SetHighlightedPane(pane int) {
	self.highlightedPane = pane
}

func (self *UiState) GetHighlightedPane() int {
	return self.highlightedPane
}

func (self *UiState) SetPopupActive(active bool) {
	self.popupActive = active
}

func (self *UiState) GetPopupActive() bool {
	return self.popupActive
}

func (self *UiState) GetActiveWidgetName() string {
	for name, widget := range self.WidgetNames {
		if widget == self.ActivePane {
			return name
		}
	}
	return ""
}
