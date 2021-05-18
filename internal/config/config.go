package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token          string `json:"token"`
	Prefix         string `json:"prefix"`
	LobbyChannel   string `json:"lobbychannel"`
	WelcomeMessage string `json:"welcomemessage"`
}

func ParseConfig(fileName string) (*Config, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	config := new(Config)

	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
