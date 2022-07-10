/*
Copyright © 2022 EGE KAAN GURKAN
*/

package cmd

import (
	"github.com/spf13/cobra"
)

// cCmd represents the c command
var cCmd = &cobra.Command{
	Use:   "c",
	Short: "Print out the current context.",
	Long:  "Print out the current context.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	//rootCmd.AddCommand(cCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
