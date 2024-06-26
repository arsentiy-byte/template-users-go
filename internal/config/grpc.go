package config

type Grpc struct {
	Host    string `yaml:"host" env-default:"127.0.0.1"`
	Port    int    `yaml:"port" env-default:"44044"`
	Timeout string `yaml:"timeout" env-default:"5s"`
}
