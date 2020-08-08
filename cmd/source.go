package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	sourceCmd.AddCommand(sourceUpdateCmd)
	rootCmd.AddCommand(sourceCmd)
}

var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "source",
}

var sourceUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updating Source Repository...")
	},
}
