package state

import (
	"sync"

	"github.com/gizak/termui/v3"
	"github.com/thegenem0/dot-mango/pkg/models"
)

var (
	uiInstance *UiState
	uiOnce     sync.Once
)

type UiState struct {
	ActiveDotfileDirChildren []models.DotfileDirectory

	// UI Management
	ActivePane  termui.Drawable
	panes       []termui.Drawable
	WidgetNames map[string]termui.Drawable
	Popup       models.Popup

	// Current States or Values
	highlightedPane     int
	selectedMangoConfig models.MangoConfig
}

func GetUiState() *UiState {
	uiOnce.Do(func() {
		uiInstance = &UiState{
			ActiveDotfileDirChildren: []models.DotfileDirectory{},
			panes:                    []termui.Drawable{},
			WidgetNames:              map[string]termui.Drawable{},
			highlightedPane:          0,
		}
	})
	return uiInstance
}

func (self *UiState) SetSelectedMangoConfig(mangoConfig models.MangoConfig) {
	self.selectedMangoConfig = mangoConfig
}

func (self *UiState) GetSelectedMangoConfig() models.MangoConfig {
	return self.selectedMangoConfig
}

func (self *UiState) SetActiveDotfileDirChildren(dotfileDirChildren []models.DotfileDirectory) {
	self.ActiveDotfileDirChildren = dotfileDirChildren
}

func (self *UiState) GetActiveDotfileDirChildren() []models.DotfileDirectory {
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
	self.Popup.IsActive = active
}

func (self *UiState) GetPopupActive() bool {
	return self.Popup.IsActive
}

func (self *UiState) SetPopupContent(content string) {
	self.Popup.Content = content
}

func (self *UiState) GetPopupContent() string {
	return self.Popup.Content
}

func (self *UiState) SetPopupType(popupType models.PopupType) {
	self.Popup.Type = popupType
}

func (self *UiState) GetPopupType() models.PopupType {
	return self.Popup.Type
}

func (self *UiState) SetPopupActions(actions []string) {
	self.Popup.Actions = actions
}

func (self *UiState) GetPopupActions() []string {
	return self.Popup.Actions
}

func (self *UiState) GetActiveWidgetName() string {
	for name, widget := range self.WidgetNames {
		if widget == self.ActivePane {
			return name
		}
	}
	return ""
}
