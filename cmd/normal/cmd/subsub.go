package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subsubOpt struct {
	flag3  string
	param3 string
}
var subsubCmd = &cobra.Command{
	Use: "subsub",
	PreRun: func(cmd *cobra.Command, args []string) {
		subsubOpt.param3 = "test333"
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("root opt: %#v\n", rootOpt)
		fmt.Printf("sub opt: %#v\n", subOpt)
		fmt.Printf("subsub opt: %#v\n", subsubOpt)
	},
}

func init() {
	subCmd.AddCommand(subsubCmd)

	subsubCmd.Flags().StringVar(&subsubOpt.flag3, "flag3", "", "")
}
