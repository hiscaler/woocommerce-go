package config

import "time"

type Config struct {
	Debug              bool          `json:"debug"` // 是否为调试模式
	WordPressAPIPrefix string        `json:"wordpress_api_prefix"`
	URL                string        `json:"url"`             // 店铺地址
	Version            string        `json:"version"`         // API 版本
	ConsumerKey        string        `json:"consumer_key"`    // Key
	ConsumerSecret     string        `json:"consumer_secret"` // Secret
	Timeout            time.Duration `json:"timeout"`         // 超时时间
	VerifySSL          bool          `json:"verify_ssl"`
	QueryStringAuth    string        `json:"query_string_auth"`
	OauthTimestamp     time.Time     `json:"oauth_timestamp"`
	EnableCache        bool          `json:"enable_cache"` // 是否激活缓存
}
