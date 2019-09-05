package configman

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

type MainConfig struct {
	DBConfig          DatabaseConfig `yaml:"database"`
	GoogleOAuthAPIKey string         `yaml:"google-oauth-api-key"`
	ListenPort        string         `yaml:"port"`
}

func ImportConfigFromFile(fileAddress string) (*MainConfig, error) {
	content, err := ioutil.ReadFile(fileAddress)
	if err != nil {
		return nil, err
	}
	var config MainConfig
	err = yaml.Unmarshal(content, &config)
	fmt.Println(config.ListenPort)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
