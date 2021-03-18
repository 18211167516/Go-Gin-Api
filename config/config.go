package config

type Config struct {
	RunMode string `mapstructure:"RUN_MODE" json:"RUN_MODE" toml:"RUN_MODE" yaml:"RUN_MODE"`
	// gorm
	Server Server `mapstructure:"server" json:"server" toml:"server" yaml:"server"`
	Mysql  Mysql  `mapstructure:"mysql" json:"app" toml:"mysql" yaml:"mysql"`
	App    App    `mapstructure:"app" json:"app" toml:"app" yaml:"app"`
	Log    Log    `mapstructure:"log" json:"log" toml:"log" yaml:"log"`
	Casbin Casbin `mapstructure:"casbin" json:"casbin" toml:"casbin" yaml:"casbin"`
}
