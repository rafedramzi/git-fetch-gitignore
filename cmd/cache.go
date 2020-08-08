package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	cacheCmd.AddCommand(cacheClearCmd)
	cacheCmd.AddCommand(cacheListCmd)

	rootCmd.AddCommand(cacheCmd)
}

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache something",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cache something ...")
	},
}

var cacheListCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache something",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cache something ...")
	},
}

var cacheClearCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache something",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cache something ...")
	},
}
