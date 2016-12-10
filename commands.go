package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config", "path to config file")
	rootCmd.AddCommand(cmdGenerate, cmdDeploy, cmdTerminate)
}

// root command (calls all other commands)
var rootCmd = &cobra.Command{
	Use:   "bora",
	Short: "bora is a simple wrapper around cloudformation. ",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var cmdGenerate = &cobra.Command{
	Use:   "generate",
	Short: "Generates a JSON or YAML template",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Creating a template")
		configReader(cfgFile)
		templateParser("exampleTemplates/sqs.yaml", os.Stdout)
	},
}

var cmdDeploy = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys a stack",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Deploying a template")
		configReader(cfgFile)
	},
}

var cmdTerminate = &cobra.Command{
	Use:   "terminate",
	Short: "Terminates a stack",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Terminating the stack")
		configReader(cfgFile)
	},
}
