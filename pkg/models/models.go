package models

type PopupType int

const (
	InfoPopup PopupType = iota
	FileOverwritePopup

	MaxPopupType
)

type MangoConfig struct {
	Name      string `yaml:"name"`
	Path      string `yaml:"path"`
	Overrides []ConfigBoundOverride
}

type ConfigBoundOverride struct {
	Path   string
	Target string
}

type Override struct {
	Name    string `yaml:"config"`
	Dotfile string `yaml:"dotfile_path"`
	Target  string `yaml:"override_target"`
}

type DotfileDirectory struct {
	Name          string
	SymlinkTarget string
	Selected      bool
	IsFolder      bool
}

type AppConfig struct {
	MangoConfigPath  string        `yaml:"mango_config_path"`
	SystemConfigPath string        `yaml:"system_config_path"`
	MangoConfigs     []MangoConfig `yaml:"dotfiles_folders"`
	Overrides        []Override    `yaml:"overrides"`
}

type Popup struct {
	IsActive bool
	Content  string
	Type     PopupType
	Actions  []string
}
