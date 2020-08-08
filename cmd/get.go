package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// TODO: Probably move this to root cmd instead

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch gitignore from defined repository",
	Run: func(cmd *cobra.Command, args []string) {
		// Check the cache dir exists
		// throw error if the cache dir not exist /perm denied with cacheflag on
		// Fetch the file
		// Throw error if it is
		// Store cache the output into cache directory
		// Print the output
		fmt.Println("Fetching repository ...")
	},
}
