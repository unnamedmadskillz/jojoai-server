package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	MongoConfig struct {
		API      string `yaml:"APIUri"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	ServerConfig struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func NewCofnig() *Config {
	c := &Config{}
	f, err := os.Open("./config/jojo.yml")
	if err != nil {
		log.Fatalf("can't read config %v, error: %v", f, err)
	}
	defer f.Close()
	yaml.NewDecoder(f).Decode(&c)
	return c
}
