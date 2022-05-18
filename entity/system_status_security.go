package entity

// SystemStatusSecurity System status security properties
type SystemStatusSecurity struct {
	SecureConnection bool `json:"secure_connection"`
	HideErrors       bool `json:"hide_errors"`
}
