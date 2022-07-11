/*
Copyright © 2022 EGE KAAN GURKAN
*/

// Package helpers contains a set of functions that are used
package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Context struct {
	Name    string        `json:"name"`
	Context ContextNested `json:"context"`
}

type ContextNested struct {
	Cluster   string `json:"cluster"`
	User      string `json:"user"`
	Namespace string `json:"namespace"`
}

type Config struct {
	PreviousContext            Context           `json:"previousContext"`
	ContextNameToLastNamespace map[string]string `json:"contextNameToLastNamespace"`
}

// GetContexts returns all contexts defined in the kube config file.
func GetContexts() []Context {

	var contexts []Context
	command := exec.Command("kubectl", "config", "view", "-o", "jsonpath='{.contexts[*]}'")
	buff := new(strings.Builder)

	command.Stdout = buff

	err := command.Run()
	if err != nil {
		logrus.Fatalf("en error occured while running the command: %e", err)
	}

	output := strings.Split(buff.String()[1:len(buff.String())-1], " ")

	for _, context := range output {
		var ctx Context
		err := json.Unmarshal([]byte(context), &ctx)
		if err != nil {
			logrus.Fatalf("an error occured while unmarshalling context: %e", err)
		}
		contexts = append(contexts, ctx)
	}

	return contexts
}

func GetAvailableContextNames() []string {

	contexts := GetContexts()
	var returnArr []string

	for _, context := range contexts {
		returnArr = append(returnArr, context.Name)
	}

	return returnArr
}

func GetAvailableContextNamesForAutocomplete(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return GetAvailableContextNames(), cobra.ShellCompDirectiveNoFileComp
}

func GetAvailableNamespaces() []string {
	output := ExecuteKubectlCommand("get", "namespace")
	splitOutput := strings.Split(output, "\n")[1:]
	var namespaces []string

	for _, out := range splitOutput {
		namespaces = append(namespaces, strings.Split(out, " ")[0])
	}

	return namespaces
}

func ExecuteKubectlCommand(args ...string) string {
	command := exec.Command("kubectl", args...)
	buff := new(strings.Builder)

	command.Stdout = buff

	err := command.Run()
	if err != nil {
		logrus.Fatalf("en error occured while running the command: %e", err)
	}

	output := buff.String()
	return output
}

func GetCurrentContext() Context {
	command := exec.Command("kubectl", "config", "view", "-o", "jsonpath='{.current-context}'")
	buff := new(strings.Builder)

	command.Stdout = buff

	err := command.Run()
	if err != nil {
		logrus.Fatalf("en error occured while running the command: %e", err)
	}

	output := buff.String()

	currentContextName := output[1 : len(output)-1]

	context, err := ContextNameToObject(currentContextName)

	if err != nil {
		logrus.Fatal(err)
	}

	return context
}

func ContextNameToObject(contextName string) (Context, error) {
	contexts := GetContexts()

	for _, context := range contexts {
		if context.Name == contextName {
			return context, nil
		}
	}

	return Context{}, fmt.Errorf("no context with name %s", contextName)
}

func SwitchContext(contextName string) (string, error) {

	previousContext := GetCurrentContext()

	command := exec.Command("kubectl", "config", "use-context", contextName)
	buff := new(strings.Builder)

	command.Stdout = buff
	err := command.Run()
	if err != nil {
		return "", err
	}

	output := buff.String()

	config := GetConfigObject()

	config.PreviousContext = previousContext
	UpdateConfigFile(config)

	return output, nil
}

func SwitchContextByObject(context Context) (string, error) {
	return SwitchContext(context.Name)
}

func SwitchNamespace(namespace string) string {
	context := GetCurrentContext()
	output := ExecuteKubectlCommand("config", "set-context", "--current", "--namespace", namespace)

	config := GetConfigObject()

	config.ContextNameToLastNamespace[context.Name] = context.Context.Namespace

	UpdateConfigFile(config)

	return output
}

func GetPreviousNamespaceOfContext(context Context) (string, error) {
	config := GetConfigObject()

	previousNamespace, exists := config.ContextNameToLastNamespace[context.Name]

	if !exists {
		return "", errors.New("please switch to a new namespace by using the 'ns' command first")
	}

	return previousNamespace, nil
}

func GetConfigFile() (*os.File, func()) {
	file, err := os.OpenFile(os.ExpandEnv("$HOME/.k8sc/config.json"), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		logrus.Fatalf("config file cannot be opened: %s", err.Error())
	}

	deferFunction := func() {
		err = file.Close()
		if err != nil {
			logrus.Fatalf("could not close the config file: %s", err.Error())
		}
	}

	return file, deferFunction
}

func GetConfigObject() Config {
	file, deferFunction := GetConfigFile()
	defer deferFunction()

	configFileContent, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.Fatalf("could not read the config file: %s", err.Error())
	}

	var config Config

	err = json.Unmarshal(configFileContent, &config)
	if err != nil {
		logrus.Fatalf("could not unmarshal JSON config file: %s", err.Error())
	}

	return config
}

func UpdateConfigFile(newConfig Config) {
	configFile, deferFunction := GetConfigFile()
	defer deferFunction()

	configJson, err := json.MarshalIndent(newConfig, "", " ")
	if err != nil {
		logrus.Fatalf("config file could not be marshaled after context switch: %s", err.Error())
	}

	err = configFile.Truncate(0)
	if err != nil {
		logrus.Fatalf("could not truncate the config file: %s", err.Error())
	}

	_, err = configFile.Seek(0, 0)
	if err != nil {
		logrus.Fatalf("could not seek to the beginning the config file: %s", err.Error())
	}

	_, err = configFile.WriteString(string(configJson))
	if err != nil {
		logrus.Fatalf("config file could not be updated after context switch: %s", err.Error())
	}
}
