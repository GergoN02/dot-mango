package ui

import (
	"fmt"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
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
	self.InfoBar.Title = "Current Config"
	self.InfoBar.Text = fmt.Sprintf(
		"::> %s -> %s\n::> %d selected items",
		self.GetSelectedMangoConfig().Name,
		self.GetSelectedMangoConfig().Path,
		self.countSelectedConfigItems(),
	)
}

func (self *View) SetConfigSelectorItems() {
	self.ConfigSelector.Rows = utils.GetConfigNames(state.GetAppState())
}
