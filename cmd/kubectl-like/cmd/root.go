package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

type RootOpt struct {
	flag1  string
	param1 string
}

func NewCmd() *cobra.Command {
	opts := &RootOpt{}

	cmd := &cobra.Command{
		Use: "root",
		Run: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(opts.Complete(cmd, args))
			cobra.CheckErr(opts.Run(cmd, args))
		},
	}
	cmd.PersistentFlags().StringVar(&opts.flag1, "flag1", "", "")

	cmd.AddCommand(NewSubCmd())
	return cmd
}

func (o *RootOpt) Complete(cmd *cobra.Command, args []string) error {
	o.param1 = "test111"
	return nil
}
func (o *RootOpt) Run(cmd *cobra.Command, args []string) error {
	return nil
}
func Execute() {
	cmd := NewCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
