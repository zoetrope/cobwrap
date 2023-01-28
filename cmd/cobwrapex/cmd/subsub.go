package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cobwrap "github.com/zoetrope/cobwrap/pkg/cobwrapex"
)

type SubSubOpt struct {
	flag3  string
	param3 string
}

func (o *SubSubOpt) Fill(cmd *cobra.Command, args []string) error {
	o.param3 = "test333"
	return nil
}

func (o *SubSubOpt) Run(cmd *cobra.Command, args []string) error {
	root := cobwrap.GetOpt[*RootOpt](cmd)
	sub := cobwrap.GetOpt[*SubOpt](cmd)
	subsub := cobwrap.GetOpt[*SubSubOpt](cmd)
	if subsub == nil {
		fmt.Println("subsub is nil")
	} else {
		fmt.Printf("subsub opt: %#v\n", subsub)
	}

	fmt.Printf("root opt: %#v\n", *root)
	fmt.Printf("sub opt: %#v\n", *sub)
	fmt.Printf("subsub opt: %#v\n", o)
	return nil
}

func NewSubSubCmd() *cobwrap.Command[*SubSubOpt] {
	cmd := &cobwrap.Command[*SubSubOpt]{
		Command: &cobra.Command{
			Use: "subsub",
		},
		Options: &SubSubOpt{},
	}
	cmd.Command.Flags().StringVar(&cmd.Options.flag3, "flag3", "", "")
	return cmd
}
