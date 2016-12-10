package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// global variables for region and stackname
var region string
var stackname string

// yamlParser parses the config YAML file with Viper
func yamlParser(conf string) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(conf)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error, can't read the config file: ", err)
	}

	// Get basic settings
	region = viper.GetString("region")
	stackname = viper.GetString("stackname")
}

// fileReader loads the meta-template
func fileReader() {}

// templateParser reads a yaml meta-template,
// and interprets it according to keys found in the configuraion
func templateParser() {} // TODO

// error checking helper function
func check(e error) {
	if e != nil {
		log.Fatal("Fatal error: ", e)
		os.Exit(1)
	}
}
