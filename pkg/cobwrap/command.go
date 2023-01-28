package cobwrap

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

type fillable interface {
	getParent() fillable
	fill(cmd *cobra.Command, args []string) error
}

type Command[T any] struct {
	Command *cobra.Command
	Fill    func(opt *T, cmd *cobra.Command, args []string) error
	RunE    func(opt *T, cmd *cobra.Command, args []string) error
	Options *T

	parent fillable
}

func (w Command[T]) getParent() fillable {
	fmt.Printf("getparent for %s(%v)\n", w.Command.Use, w.parent != nil)
	return w.parent
}

func (w Command[T]) fill(cmd *cobra.Command, args []string) error {
	err := w.Fill(w.Options, cmd, args)
	if err != nil {
		return err
	}
	fmt.Printf("add %s's option to context\n", w.Command.Use)
	fmt.Printf("new(T): %#v\n", *new(T))
	ctx := context.WithValue(cmd.Context(), *new(T), w.Options)
	cmd.SetContext(ctx)
	return nil
}

func AddCommand[P, S any](parent *Command[P], sub *Command[S]) {
	fmt.Printf("add %s to %s(%v)\n", sub.Command.Use, parent.Command.Use, parent.parent != nil)

	sub.parent = parent

	if sub.Command.RunE != nil || sub.Command.Run != nil {
		panic("error")
	}

	sub.Command.RunE = func(cmd *cobra.Command, args []string) error {
		fmt.Println("RunE")
		var parents []fillable
		for p := sub.getParent(); p != nil; p = p.getParent() {
			parents = append(parents, p)
		}
		for i := len(parents) - 1; i >= 0; i-- {
			err := parents[i].fill(cmd, args)
			if err != nil {
				return err
			}
		}
		err := sub.Fill(sub.Options, cmd, args)
		if err != nil {
			return err
		}
		return sub.RunE(sub.Options, cmd, args)
	}

	parent.Command.AddCommand(sub.Command)

	fmt.Printf("added %s to %s(%v)\n", sub.Command.Use, parent.Command.Use, parent.parent != nil)
}

func (w Command[T]) Execute() error {
	return w.Command.Execute()
}
