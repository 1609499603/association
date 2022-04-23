package config

type Server struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Casbin Casbin `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
}
