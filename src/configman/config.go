package configman

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type MainConfig struct {
	DBConfig          DatabaseConfig `yaml:"database"`
	GOAuthConfig      GoogleOAuthConfig `yaml:"google-oauth"`
	ListenPort        string         `yaml:"port"`
}

func ImportConfigFromFile(fileAddress string) (*MainConfig, error) {
	content, err := ioutil.ReadFile(fileAddress)
	if err != nil {
		return nil, err
	}
	var config MainConfig
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
