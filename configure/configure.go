package configure

import (
	"github.com/jinzhu/configor"
	"log"
	"sbe/record"
)

const (
	// Todo: make it env based
	configPath = "config/dev.yaml"
)

// New ...
func New() *record.Config {
	var config record.Config

	err := configor.New(&configor.Config{Debug: false}).Load(&config, configPath)
	if err != nil {
		log.Fatalf("Failed to load %s due to %s", configPath, err)
	}
	return &config
}
