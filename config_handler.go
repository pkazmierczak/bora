package main

import (
	"log"

	"github.com/spf13/viper"
)

// global variables for region and stackname
var region string
var stackname string

// parser parses the config YAML file with Viper
func parser(conf string) {
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error, can't read the config file: ", err)
	}

	// Get basic settings
	region = viper.GetString("region")
	stackname = viper.GetString("stackname")
}
