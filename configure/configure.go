package configure

import (
	"fmt"
	"github.com/jinzhu/configor"
	"log"
	"os"
	"sbe/record"
)

const (
	// Todo: make it env based
	configPath = "config/%s.yaml"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// New ...
func New() *record.Config {
	var config record.Config
	path := fmt.Sprintf(configPath, getenv("runenv", "dev"))
	err := configor.New(&configor.Config{Debug: false}).Load(&config, path)
	if err != nil {
		log.Printf("Failed to load %s due to %s", configPath, err)
		return nil
	}

	return &config
}
