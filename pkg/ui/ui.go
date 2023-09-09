package ui

import (
	"fmt"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/thegenem0/dot-mango/pkg/models"
	"github.com/thegenem0/dot-mango/pkg/state"
	"github.com/thegenem0/dot-mango/pkg/utils"
)

type View struct {
	*state.UiState

	Grid           *termui.Grid
	Popup          *widgets.Paragraph
	InfoBar        *widgets.Paragraph
	ConfigSelector *widgets.List
	DotfileDirs    *widgets.List
}

func (self *View) SetInfoBarText() {
	var symTarget string

	if len(self.GetActiveDotfileDirChildren()) <= 0 {
		symTarget = "None"
	} else {
		symTarget = self.GetActiveDotfileDirChildren()[0].SymlinkTarget
	}

	self.InfoBar.Title = "Current Config"
	self.InfoBar.Text = fmt.Sprintf(
		"Active Config  ::> %s -> %s\nSymlink Target ::> %s\nSelected Items ::> %d\n Active Pane   ::> %s",
		self.GetSelectedMangoConfig().Name,
		self.GetSelectedMangoConfig().Path,
		symTarget,
		self.countSelectedConfigItems(),
		self.GetActiveWidgetName(),
	)
}

func (self *View) SetConfigSelectorItems() {
	self.ConfigSelector.Rows = utils.GetConfigNames(state.GetAppState())
}

func (self *View) SetPopup() {
	self.Popup.Text = self.GetPopupContent()
	for i := models.InfoPopup; i < models.MaxPopupType; i++ {
		switch i {
		case models.InfoPopup:
			self.Popup.Title = "Info"
		case models.FileOverwritePopup:
			self.Popup.Title = "File Overwrite"
		default:
			self.Popup.Title = "Unknown"
		}
	}
}
