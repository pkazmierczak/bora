package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config type stores the configuration.
type Config struct {
	Variable, Value string
}

// Global configuration is stored here.
var C []Config

// Parser parses the config YAML file with Viper
func Parser(conf string) ([]Config, error) {
	dir, _ := os.Getwd()
	viper.AddConfigPath(dir)
	viper.SetConfigType("yaml")
	viper.SetConfigName(conf)

	// by default we back up everything
	viper.SetDefault("directories_to_skip", "none")
	viper.SetDefault("files_to_skip", "none")
	// by default we back up to Glacier
	viper.SetDefault("s3dirs", "none")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error, can't read the config file: ", err)
	}

	// put all the configuration variables into a slice
	for _, confvar := range viper.AllKeys() {
		if viper.GetString(confvar) == "" {
			log.Fatal("Fatal error, configuration variable missing: ", confvar)
		}
		C = append(C, Config{confvar, viper.GetString(confvar)})
	}
	return C, err
}

// Reader reads config variables from the global config structure.
func Reader(variable string) string {
	for i := range C {
		if C[i].Variable == variable {
			return C[i].Value
		}
	}
	log.Fatal("Fatal error, can't access configuration variable: ", variable)
	return ""
}
