package main

import (
	"fmt"
	"os"
	"veil-v2/tasks"
)

func main() {
	configPath := "config.json"

	// Check if config.json exists, if not try to load from settings.csv for migration
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("config.json not found. Please create it or migrate from settings.csv.")
		return
	}

	config, err := tasks.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	engine := tasks.NewEngine(config)
	engine.Run()
}
