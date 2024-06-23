package core

import (
	"encoding/json"
	"os"
	"path"

	"learn-go/src/types"
)
// ReadConfig reads the configuration from a file and unamarshals it into a Config struct
// returns a pointer 
func ReadConfig() (*types.Config, bool) {
	// get current file path
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, false
	}
	// creating path link to file
	filename := path.Join(rootDir,"config.json")
	// accessing file properties
	stats, err := os.Stat(filename)
	if err!= nil {
		return nil,false
	}
	// ensuring file is not empty
	if stats.Size() == 0 {
		return nil, false
	}

	// reading file contents
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, false
	}

	// setting a struct config with default server; host, port information
	config := types.Config{}

	// unamarshals JSON contents into the config struct
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, false
	}
	// returns  types.Config struct
	return &config, true
}
