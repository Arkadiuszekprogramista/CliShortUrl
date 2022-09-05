package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Redis struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"redis"`
}

func ConfigFromFile(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
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