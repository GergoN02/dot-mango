package ui

import (
	"fmt"

	"github.com/thegenem0/dot-mango/pkg/fileops"
	"github.com/thegenem0/dot-mango/pkg/models"
	"github.com/thegenem0/dot-mango/pkg/state"
)

func (self *View) SetActiveMangoConfig() {
	mangoConfigs := state.GetAppState()
	selectedIndex := self.ConfigSelector.SelectedRow
	if selectedIndex < len(mangoConfigs.GetMangoConfigs()) && selectedIndex >= 0 {
		self.SetSelectedMangoConfig(mangoConfigs.GetMangoConfigs()[selectedIndex])
	}
	configDirs := fileops.GetUserConfigDirs(self.GetSelectedMangoConfig().Path, self.GetSelectedMangoConfig().Overrides, mangoConfigs.GetSystemConfigPath())
	if len(configDirs) > 0 {
		self.SetActiveDotfileDirChildren(configDirs)
	} else {
		// TODO: Add actual logic here, this is dodgy
		self.SetActiveDotfileDirChildren([]models.DotfileDirectory{
			{
				Name:          "No dotfiles found",
				IsFolder:      false,
				SymlinkTarget: "",
				Selected:      false,
			},
		})
	}
	self.RenderActiveDotfileDirChildren()
}

func (self *View) ToggleSelectedConfigItem() {
	selectedIndex := self.DotfileDirs.SelectedRow
	if selectedIndex < len(self.ActiveDotfileDirChildren) && selectedIndex >= 0 {
		self.ActiveDotfileDirChildren[selectedIndex].Selected = !self.ActiveDotfileDirChildren[selectedIndex].Selected
		self.RenderActiveDotfileDirChildren()
	}
}

func (self *View) ToggleAllConfigItems() {
	for i := range self.ActiveDotfileDirChildren {
		self.ActiveDotfileDirChildren[i].Selected = !self.ActiveDotfileDirChildren[i].Selected
		self.RenderActiveDotfileDirChildren()
	}
}

func (self *View) RenderActiveDotfileDirChildren() {
	self.DotfileDirs.Rows = make([]string, len(self.ActiveDotfileDirChildren))
	for i, config := range self.ActiveDotfileDirChildren {
		prefix := "[ ]"
		if config.Selected {
			prefix = "[X]"
		}
		if config.IsFolder {
			icon := ""
			self.DotfileDirs.Rows[i] = fmt.Sprintf("%s %s %s -> %s", prefix, icon, config.Name, config.SymlinkTarget)
		} else {
			icon := "󰈔"
			self.DotfileDirs.Rows[i] = fmt.Sprintf("%s %s %s -> %s", prefix, icon, config.Name, config.SymlinkTarget)
		}
	}
}
