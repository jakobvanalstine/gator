package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileBase = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

var cachedConfigPath string

func Read() (*Config, error) {
	cfg := Config{}

	configPath, err := getConfigPath()
	if err != nil {
		return &cfg, err
	}

	dat, err := os.ReadFile(configPath)
	if err != nil {
		return &cfg, err
	}

	err = json.Unmarshal(dat, &cfg)
	return &cfg, err
}

func (C *Config) SetUser(name string) error {
	C.CurrentUserName = name

	dat, err := json.MarshalIndent(C, "", "  ")
	if err != nil {
		return err
	}

	_, err = C.write(dat)
	return err
}

func (C *Config) write(b []byte) (n int, err error) {
	configPath, err := getConfigPath()
	if err != nil {
		return n, err
	}

	f, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return n, err
	}

	return f.Write(b)
}

func getConfigPath() (string, error) {
	if cachedConfigPath == "" {
		dir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		cachedConfigPath = filepath.Join(dir, configFileBase)
	}

	return cachedConfigPath, nil
}
