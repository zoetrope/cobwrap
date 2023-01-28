package cmd

import (
	"github.com/spf13/cobra"
)

type SubOpt struct {
	flag2  string
	param2 string
}

func NewSubCmd() *cobra.Command {

	opts := &SubOpt{}
	cmd := &cobra.Command{
		Use: "sub",
		Run: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(opts.Complete(cmd, args))
			cobra.CheckErr(opts.Run(cmd, args))
		},
	}
	cmd.PersistentFlags().StringVar(&opts.flag2, "flag2", "", "")

	cmd.AddCommand(NewSubSubCmd())
	return cmd
}

func (o *SubOpt) Complete(cmd *cobra.Command, args []string) error {
	o.param2 = "test222"
	return nil
}
func (o *SubOpt) Run(cmd *cobra.Command, args []string) error {
	return nil
}
