package config

type System struct {
	Ip2locationPath string `mapstructure:"ip2location_path" json:"ip2location_path" yaml:"ip2location_path"`
	PageSize        int64  `mapstructure:"page_size" json:"page_size" yaml:"page_size"`
}
