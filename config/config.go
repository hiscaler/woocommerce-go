package config

import "time"

type Config struct {
	Debug                  bool          `json:"debug"`
	URL                    string        `json:"url"`
	Version                string        `json:"version"`
	ConsumerKey            string        `json:"consumer_key"`
	ConsumerSecret         string        `json:"consumer_secret"`
	AddAuthenticationToURL bool          `json:"add_authentication_to_url"`
	Timeout                time.Duration `json:"timeout"`
	VerifySSL              bool          `json:"verify_ssl"`
}
