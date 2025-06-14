package config

type JWT struct {
	ExpiresAt int    `yaml:"expiresAt"`
	Issuer    string `yaml:"issuer"`
	Key       string `yaml:"key"`
}
