package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"taskTracker/pkg/reqvalidator"
)

// ConfigPath path to config.json file
const ConfigPath = "config/config.json"

type Config struct {
	Port string `required:"true"`
}

// LoadConfig Load config metmod
func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	err = reqvalidator.Validate(&config)
	if err != nil {
		log.Printf("error loading env variables: %s\n", err.Error())
		return nil, err
	}

	return &config, nil
}
