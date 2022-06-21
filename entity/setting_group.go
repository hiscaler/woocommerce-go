package entity

// SettingGroup setting group properties
type SettingGroup struct {
	ID          string   `json:"id"`
	Label       string   `json:"label"`
	Description string   `json:"description"`
	ParentId    string   `json:"parent_id"`
	SubGroups   []string `json:"sub_groups"`
}
