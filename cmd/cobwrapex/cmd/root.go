package cmd

import (
	"os"

	"github.com/spf13/cobra"
	cobwrap "github.com/zoetrope/cobwrap/pkg/cobwrapex"
)

type RootOpt struct {
	flag1  string
	param1 string
}

func (o *RootOpt) Fill(cmd *cobra.Command, args []string) error {
	o.param1 = "test111"
	return nil
}

func (o *RootOpt) Run(cmd *cobra.Command, args []string) error {
	return nil
}

func NewCmd() *cobwrap.Command[*RootOpt] {
	cmd := &cobwrap.Command[*RootOpt]{
		Command: &cobra.Command{
			Use: "root",
		},
		Options: &RootOpt{},
	}
	cmd.Command.PersistentFlags().StringVar(&cmd.Options.flag1, "flag1", "", "")

	cobwrap.AddCommand(cmd, NewSubCmd())
	return cmd
}

func Execute() {
	cmd := NewCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
