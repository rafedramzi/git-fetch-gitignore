package cmd

import (
	"fmt"
	// homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
	"os"
)

var (
	cfgFile     string
	userLicense string
	rootCmd     = &cobra.Command{
		Use:   "get-ignore",
		Short: "Get gitignore from git repository or url",
		Long:  "Get gitignore from git repository or url",
	}
)

func init() {
	// cobra.OnInitialize()
	// rootCmd.PersistentFlags()
	// rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
}

func errExit(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
