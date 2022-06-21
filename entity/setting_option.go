package entity

// SettingOption setting option properties
type SettingOption struct {
	ID          string            `json:"id"`
	Label       string            `json:"label"`
	Description string            `json:"description"`
	Value       string            `json:"value"`
	Default     string            `json:"default"`
	Tip         string            `json:"tip"`
	PlaceHolder string            `json:"place_holder"`
	Type        string            `json:"type"`
	Options     map[string]string `json:"options"`
	GroupId     string            `json:"group_id"`
}
