package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// config type stores the configuration.
type config struct {
	Variable, Value string
}

// Global configuration is stored here.
var c []config

// parser parses the config YAML file with Viper
func parser(conf string) ([]config, error) {
	dir, _ := os.Getwd()
	viper.AddConfigPath(dir)
	viper.SetConfigType("yaml")
	viper.SetConfigName(conf)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error, can't read the config file: ", err)
	}

	// put all the configuration variables into a slice
	for _, confvar := range viper.AllKeys() {
		if viper.GetString(confvar) == "" {
			log.Fatal("Fatal error, cannot find value for variable: ", confvar)
		}
		c = append(c, config{confvar, viper.GetString(confvar)})
	}
	return c, err
}

// reader reads config variables from the global config structure.
func reader(variable string) string {
	for i := range c {
		if c[i].Variable == variable {
			return c[i].Value
		}
	}
	log.Fatal("Fatal error, can't access configuration variable: ", variable)
	return ""
}
