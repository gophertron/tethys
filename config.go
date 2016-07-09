package main

type Config struct {
	RepoRoot string
}

var config *Config = nil

func SetConfig(cfg *Config) {
	config = cfg
}

func GetConfig() *Config {
	return config
}
