package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var tplFile string

func init() {
	generateCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config", "path to config file")
	generateCmd.PersistentFlags().StringVarP(&tplFile, "template", "t", "template", "path to template file")
	deployCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config", "path to config file")
	deployCmd.PersistentFlags().StringVarP(&tplFile, "template", "t", "template", "path to template file")
	terminateCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config", "path to config file")
	rootCmd.AddCommand(generateCmd, deployCmd, terminateCmd)
}

// root command (calls all other commands)
var rootCmd = &cobra.Command{
	Use:   "bora",
	Short: "bora is a simple wrapper around cloudformation. ",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a JSON or YAML template",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Creating a template")
		configReader(cfgFile)
		templateParser(tplFile, os.Stdout)
	},
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys a stack",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Deploying a template")
		configReader(cfgFile)
	},
}

var terminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: "Terminates a stack",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Terminating the stack")
		configReader(cfgFile)
	},
}
