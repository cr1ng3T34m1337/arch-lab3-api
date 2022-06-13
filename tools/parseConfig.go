package tools

import (
	"encoding/json"
	"os"
)

type Config struct {
	Db struct {
		DbName   string `json:"name"`
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
	} `json:"db"`
	ApiUrl string `json:"apiUrl"`
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
