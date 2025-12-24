package tasks

import (
	"encoding/json"
	"os"
)

type TaskConfig struct {
	Username              string   `json:"username"`
	Password              string   `json:"password"`
	Term                  string   `json:"term"`
	Subject               string   `json:"subject"`
	Mode                  string   `json:"mode"`
	CRNs                  []string `json:"crns"`
	WebhookURL            string   `json:"webhook_url"`
	RegistrationTime      string   `json:"registration_time"`
	SavedRegistrationTime string   `json:"saved_registration_time,omitempty"`
}

type Config struct {
	Tasks []TaskConfig `json:"tasks"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func SaveConfig(path string, config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
