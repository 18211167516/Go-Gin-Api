package config

type Config struct {
	Server Server `mapstructure:"server" json:"server" toml:"server" yaml:"server"`
	App    App    `mapstructure:"app" json:"app" toml:"app" yaml:"app"`
	Log    Log    `mapstructure:"log" json:"log" toml:"log" yaml:"log"`
}
