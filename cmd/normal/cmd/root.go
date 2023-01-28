package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootOpt struct {
	flag1  string
	param1 string
}

var rootCmd = &cobra.Command{
	Use: "normal",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		rootOpt.param1 = "test111"
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&rootOpt.flag1, "flag1", "", "")
}
