package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App      App      `yaml:"app"`
	Database Database `yaml:"database"`
	Allows   Allows   `yaml:"allows"`
}

type App struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type Database struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Name string `yaml:"name"`
}

type Allows struct {
	Methods []string `yaml:"methods"`
	Origins []string `yaml:"origins"`
	Headers []string `yaml:"headers"`
}

func InitConfig() *Config {
	var configs Config
	file_name, _ := filepath.Abs("./config.yaml")
	yaml_file, _ := os.ReadFile(file_name)
	yaml.Unmarshal(yaml_file, &configs)
	return &configs
}
