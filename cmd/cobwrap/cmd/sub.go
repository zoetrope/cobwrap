package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zoetrope/cobwrap/pkg/cobwrap"
)

type SubOpt struct {
	flag2  string
	param2 string
}

func NewSubCmd() *cobwrap.Command[SubOpt] {
	cmd := &cobwrap.Command[SubOpt]{
		Command: &cobra.Command{
			Use: "sub",
		},
		Fill: func(opt *SubOpt, cmd *cobra.Command, args []string) error {
			fmt.Println("sub.Fill")
			opt.param2 = "test222"
			return nil
		},
		Options: &SubOpt{},
	}
	cmd.Command.PersistentFlags().StringVar(&cmd.Options.flag2, "flag2", "", "")

	cobwrap.AddCommand(cmd, NewSubSubCmd())
	return cmd
}
