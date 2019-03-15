package conf

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type ConfigManager struct {
	file string
}

func NewConfig(file string) *ConfigManager {
	return &ConfigManager{file}
}

func (cm *ConfigManager) Load() (*Config, error) {
	data, err := ioutil.ReadFile(cm.file)
	if err != nil {
		return nil, err
	}

	var config *Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

type Config struct {
	Debug bool `yaml:"debug"`
	DB    DB   `yaml:"db"`
	Host  Host `yaml:"host"`
}

type DB struct {
	UserName     string `yaml:"username"`
	UserPassword string `yaml:"userpassword"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Database     string `yaml:"database"`
}

type Host struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}
