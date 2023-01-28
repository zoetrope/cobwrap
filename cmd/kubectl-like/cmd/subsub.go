package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type SubSubOpt struct {
	flag3  string
	param3 string
}

func NewSubSubCmd() *cobra.Command {
	opts := &SubSubOpt{}
	cmd := &cobra.Command{
		Use: "subsub",
		Run: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(opts.Complete(cmd, args))
			cobra.CheckErr(opts.Run(cmd, args))
		},
	}

	cmd.Flags().StringVar(&opts.flag3, "flag3", "", "")

	return cmd
}

func (o *SubSubOpt) Complete(cmd *cobra.Command, args []string) error {
	o.param3 = "test333"
	return nil
}
func (o *SubSubOpt) Run(cmd *cobra.Command, args []string) error {
	//fmt.Printf("root opt: %#v\n", RootOpt{})
	//fmt.Printf("sub opt: %#v\n", SubSubOpt{})
	fmt.Printf("subsub opt: %#v\n", o)
	return nil
}
