package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server,omitempty"`
	Postgres PostgresConfig `yaml:"postgres,omitempty"`
}

type ServerConfig struct {
	AppVersion string `yaml:"app_version,omitempty"`
	Port       int    `yaml:"port,omitempty"`
	Mode       string `yaml:"mode,omitempty"`
	JwtKey     string `yaml:"jwtkey,omitempty"`
}

type PostgresConfig struct {
	Host     string `yaml:"host,omitempty"`
	Port     int    `yaml:"port,omitempty"`
	Dbname   string `yaml:"dbname,omitempty"`
	User     string `yaml:"user,omitempty"`
	Password string `yaml:"password,omitempty"`
}

func ReadConfig(path string) *Config {
	var config Config

	fb, err := os.ReadFile(path)
	if err != nil {
		log.Printf("ReadFile: %v", err)
		panic(err)
	}

	err = yaml.Unmarshal(fb, &config)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
		panic(err)
	}

	return &config
}
