package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Redis struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"redis"`
}

func ConfigFromFile(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil

}
