package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var stack string
var tplFile string

func init() {
	generateCmd.Flags().StringVarP(&cfgFile, "config", "c", "config", "path to config file")
	generateCmd.Flags().StringVarP(&tplFile, "template", "t", "template", "path to template file")
	deployCmd.Flags().StringVarP(&cfgFile, "config", "c", "config", "path to config file")
	deployCmd.Flags().StringVarP(&tplFile, "template", "t", "template", "path to template file")
	terminateCmd.Flags().StringVarP(&stack, "stack", "s", "stack", "name of the stack to terminate")
	rootCmd.AddCommand(generateCmd, deployCmd, terminateCmd)
	viper.SetDefault("author", "Piotr Kazmierczak <me@piotrkazmierczak.com>")
	viper.SetDefault("license", "MIT")
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
		log.Println("Generating template")
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
		log.Println("Terminating the stack", stack)
	},
}
