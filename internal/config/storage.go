package config

type Storage struct {
	Host     string `yaml:"host" env-default:"127.0.0.1"`
	Port     int    `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Database string `yaml:"database" env-required:"true"`
}
