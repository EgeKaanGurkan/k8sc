/*
Copyright © 2022 EGE KAAN GURKAN
*/

package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"k8sc/helpers"
	"os"
)

// csCmd represents the cs command
var csCmd = &cobra.Command{
	Use:               "cs",
	Aliases:           []string{"sc"},
	Short:             "Switch context [c]ontext [s]witch",
	Long:              "Switch context [c]ontext [s]witch",
	ValidArgsFunction: helpers.GetAvailableContextNamesForAutocomplete,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println(fmt.Errorf("please provide a context"))
			os.Exit(1)
		}

		contextName := args[0]

		contextSwitchOutput, err := helpers.SwitchContext(contextName)
		if err != nil {
			logrus.Fatalf("error switching to context %s: %e", contextName, err)
		}

		fmt.Print(contextSwitchOutput)

		context := helpers.GetCurrentContext()
		fmt.Printf("Current namespace: %q\n", context.Context.Namespace)

	},
}

func init() {
	rootCmd.AddCommand(csCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// csCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// csCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
