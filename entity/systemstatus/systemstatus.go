package systemstatus

type SystemStatus struct {
	Environment   Environment `json:"environment"`
	Database      Database    `json:"database"`
	ActivePlugins []string    `json:"active_plugins"`
	Theme         Theme       `json:"theme"`
	Settings      Setting     `json:"settings"`
	Security      Security    `json:"security"`
	Pages         []string    `json:"pages"`
}
