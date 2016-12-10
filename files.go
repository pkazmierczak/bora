package main

import (
	"io"
	"io/ioutil"
	"log"
	"text/template"

	"github.com/spf13/viper"
)

// global variables for region and stackname
var region string
var stackname string

// cfvars go into a global map
var cfvars map[string]interface{}

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

	// Get all the values of variables under CF
	// and put them into a map
	cfvars = make(map[string]interface{})
	for _, confvar := range viper.Sub("CF").AllKeys() {
		if confvar != "" {
			confval := viper.Get("CF." + confvar)
			cfvars[confvar] = confval
		}
	}
}

func templateParser(filename string, wr io.Writer) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading the template file: ", err)
	}

	t, err := template.New("template").Parse(string(b))
	if err != nil {
		log.Fatal("Error parsing the template file: ", err)
	}
	t.Execute(wr, cfvars)
}
