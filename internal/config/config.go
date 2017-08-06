package config

type Config struct {
	ListId string `toml:"list-id"`
	ApiKey string `toml:"api-key"`
	Origin string `toml:"origin"`
}
