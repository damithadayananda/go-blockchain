package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var AppConfig config

type config struct {
	Complexity int8 `yaml:"complexity"`
}

func InitConfig() {
	b, e := ioutil.ReadFile("./config/config.yml")
	if e != nil {
		panic(e)
	}
	if err := yaml.Unmarshal(b, &AppConfig); err != nil {
		panic(err)
	}
}
