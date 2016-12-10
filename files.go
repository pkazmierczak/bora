package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"

	yaml "gopkg.in/yaml.v2"

	"github.com/spf13/viper"
)

// global variables for region and stackname
var region string
var stackname string

// configReader parses the config YAML file with Viper
func configReader(conf string) {
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

// templateReader loads the meta-template
// returns a yaml unmarshalled into an empty interface
func templateReader(filename string) map[string]interface{} {
	// we unmarhshall into an interface
	var i interface{}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading the template file: ", err)
	}
	yaml.Unmarshal([]byte(yamlFile), &i)
	m := i.(map[string]interface{}) // a bit hacky, I know...
	return m
}

}
