// File: internal/config/config.go

package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port            string `json:"port"`
	TokenSecret     string `json:"token_secret"`
	MinDepositAmount float64 `json:"min_deposit_amount"`
}

func Load() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	// Set default values if not specified
	if config.Port == "" {
		config.Port = "8080"
	}
	if config.MinDepositAmount == 0 {
		config.MinDepositAmount = 5
	}

	return &config, nil
}