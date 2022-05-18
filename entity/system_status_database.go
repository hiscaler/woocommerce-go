package entity

// SystemStatusDatabase System status database properties
type SystemStatusDatabase struct {
	WCDatabaseVersion    string   `json:"wc_database_version"`
	DatabasePrefix       string   `json:"database_prefix"`
	MaxmindGEOIPDatabase string   `json:"maxmind_geoip_database"`
	DatabaseTables       []string `json:"database_tables"`
}
