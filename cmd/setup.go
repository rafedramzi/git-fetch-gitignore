package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var setupInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup something",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Setup init something ...")
	},
}

var setupClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Setup clear",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Setup clear something ...")
	},
}

var setupCmd = &cobra.Command{
	Use:   "setupCmd",
	Short: "Setup something",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Setup something ...")
	},
}

func init() {
	setupCmd.AddCommand(setupInitCmd)
	setupCmd.AddCommand(setupClearCmd)

	rootCmd.AddCommand(setupCmd)
}
