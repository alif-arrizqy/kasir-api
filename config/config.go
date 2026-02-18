package config

type Config struct {
	Host   string `mapstructure:"HOST"`
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DBConn"`
	APIKey string `mapstructure:"API_KEY"`
}