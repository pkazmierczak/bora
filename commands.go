package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var cfgFile string
var tplFile string

func init() {
	generateCmd.Flags().StringVarP(&cfgFile, "config", "c", "config", "path to config file")
	generateCmd.Flags().StringVarP(&tplFile, "template", "t", "template", "path to template file")
	deployCmd.Flags().StringVarP(&cfgFile, "config", "c", "config", "path to config file")
	deployCmd.Flags().StringVarP(&tplFile, "template", "t", "template", "path to template file")
	terminateCmd.Flags().StringVarP(&cfgFile, "config", "c", "config", "path to config file")
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
		configReader(cfgFile)
		log.Println("Generating a template for", stackname)
		tpl := templateParser(tplFile)
		fmt.Println(tpl)
	},
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys a stack",
	Run: func(cmd *cobra.Command, args []string) {
		configReader(cfgFile)
		log.Println("Deploying a template for", stackname)
		tpl := templateParser(tplFile)
		_, sess := awsSession()
		deployStack(tpl, sess)
	},
}

var terminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: "Terminates a stack",
	Run: func(cmd *cobra.Command, args []string) {
		configReader(cfgFile)
		log.Println("Terminating stack", stackname)
		_, sess := awsSession()
		terminateStack(sess)
	},
}
