package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token          string `json:"token"`
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	ServerID       string `json:"server_id"`
	Prefix         string `json:"prefix"`
	LobbyChannel   string `json:"lobbychannel"`
	InviteLink     string `json:"invite_link"`
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
