/*
Copyright © 2022 EGE KAAN GURKAN
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8sc/helpers"
	"os"
	"strings"
)

// gpCmd represents the gp command
var gpCmd = &cobra.Command{
	Use:     "gp",
	Short:   "Get pods",
	Aliases: []string{"pg"},
	Long:    `Get pods with different operations.`,
	Run: func(cmd *cobra.Command, args []string) {

		command := "get pod"

		nodeName, err := cmd.Flags().GetString("node")
		if err != nil {
			fmt.Printf("There was an error retrieving the flag value of 'node': %s\n", err.Error())
			os.Exit(1)
		}

		wide, err := cmd.Flags().GetBool("wide")
		if err != nil {
			fmt.Printf("There was an error retrieving the flag value of 'wide': %s\n", err.Error())
			os.Exit(1)
		}

		if wide {
			command += " -o wide"
		}

		if nodeName != "" {
			command += fmt.Sprintf(" --field-selector spec.nodeName=%s", nodeName)
		}

		commandSplit := strings.Split(command, " ")

		output := helpers.ExecuteKubectlCommand(commandSplit...)

		fmt.Print(output)
	},
}

// getAvailableNodesForAutocomplete returns the available nodes as a string array and also a cobra autocomplete
// directive stating that we don't want file autocompletion for this context.
func getAvailableNodesForAutocomplete(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return helpers.GetAvailableNodes(), cobra.ShellCompDirectiveNoFileComp
}

func init() {
	rootCmd.AddCommand(gpCmd)

	gpCmd.Flags().StringP("node", "n", "", "get pods within a node")
	err := gpCmd.RegisterFlagCompletionFunc("node", getAvailableNodesForAutocomplete)
	if err != nil {
		fmt.Printf("could not register the completion function for the 'node' flag. Please use kubectl or k8sc" +
			"to manually retrieve node names.\n")
	}

	gpCmd.Flags().BoolP("wide", "w", false, "retrieve wide output")
}
