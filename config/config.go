package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

var AppConfig config

type config struct {
	Complexity               int8     `yaml:"complexity"`
	MaxTransactionsPerBlock  int      `yaml:"maxTransactionsPerBlock"`
	Host                     string   `yaml:"host"`
	Port                     int      `yaml:"port"`
	NodeDistributionTimeOut  int      `yaml:"nodeDistributionTimeOut"`
	BlockDistributionTimeOut int      `yaml:"blockDistributionTimeout"`
	KnownNodes               []string `yaml:"knownNodes"`
}

func InitConfig() {
	configPath := os.Getenv("CONFIG_PATH")
	b, e := os.ReadFile(configPath)
	if e != nil {
		panic(e)
	}
	if err := yaml.Unmarshal(b, &AppConfig); err != nil {
		panic(err)
	}
}
