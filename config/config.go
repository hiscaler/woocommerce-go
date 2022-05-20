package config

import "time"

type Config struct {
	Debug                  bool          `json:"debug"`                     // 是否为调试模式
	URL                    string        `json:"url"`                       // 店铺地址
	Version                string        `json:"version"`                   // API 版本
	ConsumerKey            string        `json:"consumer_key"`              // Consumer Key
	ConsumerSecret         string        `json:"consumer_secret"`           // Consumer Secret
	AddAuthenticationToURL bool          `json:"add_authentication_to_url"` // 是否将认证信息附加到 URL 中
	Timeout                time.Duration `json:"timeout"`                   // 超时时间（秒）
	VerifySSL              bool          `json:"verify_ssl"`                // 是否验证 SSL
}
