package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subOpt struct {
	flag2  string
	param2 string
}

var subCmd = &cobra.Command{
	Use: "sub",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("persistent sub")
		subOpt.param2 = "test222"
	},
}

func init() {
	rootCmd.AddCommand(subCmd)

	subCmd.PersistentFlags().StringVar(&subOpt.flag2, "flag2", "", "")
}
