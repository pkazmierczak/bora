package main

import (
	"log"

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
	Run:   generateRun,
}

func generateRun(cmd *cobra.Command, args []string) {
	log.Println("Creating a template")
	yamlParser(cfgFile)
}

var cmdDeploy = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys a stack",
	Run:   deployRun,
}

func deployRun(cmd *cobra.Command, args []string) {
	log.Println("Deploying a template")
	yamlParser(cfgFile)
}

var cmdTerminate = &cobra.Command{
	Use:   "terminate",
	Short: "Terminates a stack",
	Run:   terminateRun,
}

func terminateRun(cmd *cobra.Command, args []string) {
	log.Println("Terminating the stack")
	yamlParser(cfgFile)
}
