package config

import "time"

type Config struct {
	Debug          bool          `json:"debug"`             // 是否为调试模式
	URL            string        `json:"url"`               // 店铺地址
	Version        string        `json:"version"`           // API 版本
	ConsumerKey    string        `json:"consumer_key"`      // Consumer Key
	ConsumerSecret string        `json:"consumer_secret"`   // Consumer Secret
	UseAuthInQuery bool          `json:"use_auth_in_query"` // 是否将认证内容放在 URL.Query 中
	Timeout        time.Duration `json:"timeout"`           // 超时时间（秒）
	VerifySSL      bool          `json:"verify_ssl"`        // 是否验证 SSL
	EnableCache    bool          `json:"enable_cache"`      // 是否激活缓存
}
