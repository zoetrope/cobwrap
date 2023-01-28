package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zoetrope/cobwrap/pkg/cobwrap"
)

type RootOpt struct {
	flag1  string
	param1 string
}

func NewCmd() *cobwrap.Command[RootOpt] {
	cmd := &cobwrap.Command[RootOpt]{
		Command: &cobra.Command{
			Use: "root",
		},
		Fill: func(opt *RootOpt, cmd *cobra.Command, args []string) error {
			fmt.Println("root.Fill")
			opt.param1 = "test111"
			return nil
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
