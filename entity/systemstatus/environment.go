package systemstatus

// Environment System status - Environment properties
type Environment struct {
	HomeURL                string `json:"home_url"`
	SiteURL                string `json:"site_url"`
	WCVersion              string `json:"wc_version"`
	LogDirectory           string `json:"log_directory"`
	LogDirectoryWritable   bool   `json:"log_directory_writable"`
	WPVersion              string `json:"wp_version"`
	WPMultisite            bool   `json:"wp_multisite"`
	WPMemoryLimit          int    `json:"wp_memory_limit"`
	WPDebugMode            bool   `json:"wp_debug_mode"`
	WPCron                 bool   `json:"wp_cron"`
	Language               string `json:"language"`
	ServerInfo             string `json:"server_info"`
	PHPVersion             string `json:"php_version"`
	PHPPostMaxSize         int    `json:"php_post_max_size"`
	PHPMaxExecutionTime    int    `json:"php_max_execution_time"`
	PHPMaxInputVars        int    `json:"php_max_input_vars"`
	CURLVersion            string `json:"curl_version"`
	SuhosinInstalled       bool   `json:"suhosin_installed"`
	MaxUploadSize          int    `json:"max_upload_size"`
	MySQLVersion           string `json:"my_sql_version"`
	DefaultTimezone        string `json:"default_timezone"`
	FSockOpenOrCurlEnabled bool   `json:"fsockopen_or_curl_enabled"`
	SOAPClientEnabled      bool   `json:"soap_client_enabled"`
	GzipEnabled            bool   `json:"gzip_enabled"`
	MbStringEnabled        bool   `json:"mbstring_enabled"`
	RemotePostSuccessful   bool   `json:"remote_post_successful"`
	RemotePostResponse     string `json:"remote_post_response"`
	RemoteGetSuccessful    bool   `json:"remote_get_successful"`
	RemoteGetResponse      string `json:"remote_get_response"`
}
