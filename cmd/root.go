/*
Copyright © 2022 EGE KAAN GURKAN
*/

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"k8sc/helpers"
	"os"
	"os/exec"
)

var Version = "0.0.1"
var Verbose bool
var displayVersion bool

var ConfigDir = os.ExpandEnv("$HOME/.k8sc")
var ConfigFile = "config.json"
var ConfigPath = fmt.Sprintf("%s/%s", ConfigDir, ConfigFile)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "k8sc",
	Short: "A convenience script for kubectl",
	Long: `The k8sc script aims to ease the use if kubectl, especially for commands that are used daily
which take too long to type out and those that are overall convenient, like switching to a previous namespace.`,
	PersistentPreRun: initialSetup,
	Run: func(cmd *cobra.Command, args []string) {

		if displayVersion {
			fmt.Println(fmt.Sprintf("Version: %s", Version))
			os.Exit(0)
		}

		currentContext := helpers.GetCurrentContext()

		if currentContext.Context.Namespace == "" {
			currentContext.Context.Namespace = "default"
		}

		fmt.Println("Current Context")
		fmt.Printf("Name: %s\n", currentContext.Name)
		fmt.Printf("Namespace: %s\n", currentContext.Context.Namespace)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initialSetup(cmd *cobra.Command, args []string) {

	// set logging level
	if Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// check if kubectl exists
	_, err := exec.LookPath("kubectl")
	if err != nil {
		logrus.Fatalln("kubectl not present")
	}

	// Create the config file and directory
	err = os.Mkdir(ConfigDir, os.ModePerm)
	if err != nil {
		logrus.Debugf("problem creating the config directory: %s", err.Error())
	}

	configFile, err := os.OpenFile(fmt.Sprintf("%s", ConfigPath), os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		logrus.Fatalf("there was a problem opening or creating the config file: %s\n", err.Error())
	}

	fileContents, err := ioutil.ReadAll(configFile)
	if err != nil {
		logrus.Debugf("could not read the config file contents: %s", err.Error())
	}

	if len(fileContents) == 0 {
		var config helpers.Config

		config.ContextNameToLastNamespace = make(map[string]string)

		jsonData, err := json.MarshalIndent(config, "", " ")
		if err != nil {
			logrus.Fatalf("could not initiate the config file: %s", err.Error())
		}

		_, err = configFile.WriteString(string(jsonData))
		if err != nil {
			logrus.Fatalf("could not write to the config file: %s", err.Error())
		}
	}

	if Verbose {
		fmt.Println("Config File Contents")
		fmt.Printf("%s\n", fileContents)
	}

	err = configFile.Close()
	if err != nil {
		logrus.Fatalf("could not close the config file: %s", err.Error())
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.k8sc.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "V", false, "verbose mode")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&displayVersion, "version", "v", false, "display current version")
}
