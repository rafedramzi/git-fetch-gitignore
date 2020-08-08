package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateRepositoryCmd)
}

var updateRepositoryCmd = &cobra.Command{
	Use:   "update-repository",
	Short: "Update repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updating Repository...")
	},
}
