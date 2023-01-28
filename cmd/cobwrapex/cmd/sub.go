package cmd

import (
	"github.com/spf13/cobra"
	cobwrap "github.com/zoetrope/cobwrap/pkg/cobwrapex"
)

type SubOpt struct {
	flag2  string
	param2 string
}

func (o *SubOpt) Fill(cmd *cobra.Command, args []string) error {
	o.param2 = "test222"
	return nil
}

func (o *SubOpt) Run(cmd *cobra.Command, args []string) error {
	return nil
}

func NewSubCmd() *cobwrap.Command[*SubOpt] {

	cmd := &cobwrap.Command[*SubOpt]{
		Command: &cobra.Command{
			Use: "sub",
		},
		Options: &SubOpt{},
	}
	cmd.Command.PersistentFlags().StringVar(&cmd.Options.flag2, "flag2", "", "")

	cobwrap.AddCommand(cmd, NewSubSubCmd())
	return cmd
}
