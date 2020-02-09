package config

type AppCfg struct {
	Mysql mysql `yaml:"mysql"`
	Redis redis `yaml:"redis"`
	Mongo mongo `yaml:"mongo"`
}

type mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type mongo struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	DB         string `yaml:"db"`
	Timeout    int    `yaml:"timeout"`
	AuthSource string `yaml:"auth_source"`
}
