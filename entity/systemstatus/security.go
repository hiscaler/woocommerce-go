package systemstatus

// Security System status - Security properties
type Security struct {
	SecureConnection bool `json:"secure_connection"`
	HideErrors       bool `json:"hide_errors"`
}
