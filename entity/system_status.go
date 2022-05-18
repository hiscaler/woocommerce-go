package entity

type SystemStatus struct {
	Environment   SystemStatusEnvironment `json:"environment"`
	Database      SystemStatusDatabase    `json:"database"`
	ActivePlugins []string                `json:"active_plugins"`
	Theme         SystemStatusTheme       `json:"theme"`
	Settings      SystemStatusSetting     `json:"settings"`
	Security      SystemStatusSecurity    `json:"security"`
	Pages         []string                `json:"pages"`
}
