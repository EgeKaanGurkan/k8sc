/*
Copyright © 2022 EGE KAAN GURKAN
*/

package cmd

import (
	"fmt"
	"k8sc/helpers"
	"os"

	"github.com/spf13/cobra"
)

// nsCmd represents the ns command
var nsCmd = &cobra.Command{
	Use:               "ns",
	Aliases:           []string{"sn"},
	Short:             "Switch namespace",
	Long:              `Switch namespace`,
	ValidArgsFunction: getAvailableNamespacesForAutoComplete,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println(fmt.Errorf("please provide a namespace"))
			os.Exit(1)
		}

		selectedNamespace := args[0]

		output := helpers.SwitchNamespace(selectedNamespace)

		fmt.Println(output)
	},
}

func getAvailableNamespacesForAutoComplete(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return helpers.GetAvailableNamespaces(), cobra.ShellCompDirectiveNoFileComp
}

func init() {
	rootCmd.AddCommand(nsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
