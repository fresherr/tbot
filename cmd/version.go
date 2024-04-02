/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	appVersion = "0.0.1"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Current app version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(appVersion)
	},
}


func init() {
	rootCmd.AddCommand(versionCmd)
}
