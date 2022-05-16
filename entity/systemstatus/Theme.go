package systemstatus

// Theme System status - Theme properties
type Theme struct {
	Name                  string   `json:"name"`
	Version               string   `json:"version"`
	VersionLatest         string   `json:"version_latest"`
	AuthorURL             string   `json:"author_url"`
	IsChildTheme          bool     `json:"is_child_theme"`
	HasWooCommerceSupport bool     `json:"has_woo_commerce_support"`
	HasWooCommerceFile    bool     `json:"has_woo_commerce_file"`
	HasOutdatedTemplates  bool     `json:"has_outdated_templates"`
	Overrides             []string `json:"overrides"`
	ParentName            string   `json:"parent_name"`
	ParentVersion         string   `json:"parent_version"`
	ParentAuthorURL       string   `json:"parent_author_url"`
}
