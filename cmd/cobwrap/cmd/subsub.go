package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zoetrope/cobwrap/pkg/cobwrap"
)

type SubSubOpt struct {
	flag3  string
	param3 string
}

func NewSubSubCmd() *cobwrap.Command[SubSubOpt] {
	cmd := &cobwrap.Command[SubSubOpt]{
		Command: &cobra.Command{
			Use: "subsub",
		},
		Fill: func(opt *SubSubOpt, cmd *cobra.Command, args []string) error {
			fmt.Println("subsub.Fill")
			opt.param3 = "test333"
			return nil
		},
		Options: &SubSubOpt{},
		RunE: func(opt *SubSubOpt, cmd *cobra.Command, args []string) error {
			root := cmd.Context().Value(RootOpt{}).(*RootOpt)
			sub := cmd.Context().Value(SubOpt{}).(*SubOpt)
			fmt.Printf("root opt: %#v\n", *root)
			fmt.Printf("sub opt: %#v\n", *sub)
			fmt.Printf("subsub opt: %#v\n", *opt)
			return nil
		},
	}
	cmd.Command.Flags().StringVar(&cmd.Options.flag3, "flag3", "", "")
	return cmd
}
