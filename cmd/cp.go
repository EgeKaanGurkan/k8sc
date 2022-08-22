/*
Copyright © 2022 EGE KAAN GURKAN
*/

package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"k8sc/helpers"
)

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:     "cp",
	Aliases: []string{"pc"},
	Short:   "Switch to the previous context. [c]ontext [p]revious",
	Long:    `Switch to the previous context. [c]ontext [p]revious`,
	Run: func(cmd *cobra.Command, args []string) {

		config := helpers.GetConfigObject()

		// Current here, will be previous after context switch
		previousContext := helpers.GetCurrentContext()

		switchToContext := config.PreviousContext

		// There hasn't been a context switch before.
		if switchToContext.Name == "" {
			fmt.Println(fmt.Errorf("please switch to a new context by using the 'cs' command first"))
			return
		}

		output, err := helpers.SwitchContextByObject(config.PreviousContext)
		if err != nil {
			logrus.Fatalf("could not switch context: %s", err.Error())
		}

		config.PreviousContext = previousContext

		helpers.UpdateConfigFile(config)

		fmt.Print(output)
	},
}

func init() {
	rootCmd.AddCommand(cpCmd)
}
