package entity

// SystemStatusTool system status tool properties
type SystemStatusTool struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Action      string `json:"action"`
	Description string `json:"description"`
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	Confirm     bool   `json:"confirm"`
}
