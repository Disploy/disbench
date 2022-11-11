package cmd

import (
	"github.com/spf13/cobra"
)

var internalCmd = &cobra.Command{
	Use:   "internal",
	Short: "Tools for internally testing Disploy.",
	Long:  `These commands are for internal use only. They are not intended for testing your own bots made with Disploy.`,
}

func init() {
	rootCmd.AddCommand(internalCmd)
}
