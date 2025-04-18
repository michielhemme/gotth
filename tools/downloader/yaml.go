package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

func loadConfig(filename string) (DownloadConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var cfg DownloadConfig
	err = yaml.Unmarshal(data, &cfg)
	return cfg, err
}
