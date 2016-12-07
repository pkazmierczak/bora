package main

import (
	"log"

	"github.com/spf13/viper"
)

// global variables for region and stackname
var region string
var stackname string

// yamlParser parses the config YAML file with Viper
func yamlParser(conf string) {
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error, can't read the config file: ", err)
	}

	// Get basic settings
	region = viper.GetString("region")
	stackname = viper.GetString("stackname")
}

// templateParser reads a yaml meta-template,
// and interprets it according to keys found in the configuraion
func templateParser() {} // TODO
