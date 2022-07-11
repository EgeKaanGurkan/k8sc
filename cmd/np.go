/*
Copyright © 2022 EGE KAAN GURKAN
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8sc/helpers"
	"os"
)

// npCmd represents the np command
var npCmd = &cobra.Command{
	Use:     "np",
	Aliases: []string{"pn"},
	Short:   "Switches to the previous namespace for this context",
	Long:    "Switches to the previous namespace for this context",
	Run: func(cmd *cobra.Command, args []string) {
		config := helpers.GetConfigObject()
		context := helpers.GetCurrentContext()

		previousNamespace := context.Context.Namespace
		switchToNamespace, err := helpers.GetPreviousNamespaceOfContext(context)
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		}

		output := helpers.SwitchNamespace(switchToNamespace)

		config.ContextNameToLastNamespace[context.Name] = previousNamespace

		helpers.UpdateConfigFile(config)

		fmt.Print(output)
	},
}

func init() {
	rootCmd.AddCommand(npCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// npCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// npCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
