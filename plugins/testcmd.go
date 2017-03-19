package main

import (
	"context"
	"fmt"
	"io"

	"github.com/vladimirvivien/gosh/api"
)

type helloCmd string

func (t helloCmd) Name() string      { return string(t) }
func (t helloCmd) Usage() string     { return `hello` }
func (t helloCmd) ShortDesc() string { return `prints greeting "hello there"` }
func (t helloCmd) LongDesc() string  { return t.ShortDesc() }
func (t helloCmd) Exec(ctx context.Context, args []string) (context.Context, error) {
	out := ctx.Value("gosh.stdout").(io.Writer)
	fmt.Fprintln(out, "hello there")
	return ctx, nil
}

type goodbyeCmd string

func (t goodbyeCmd) Name() string      { return string(t) }
func (t goodbyeCmd) Usage() string     { return t.Name() }
func (t goodbyeCmd) ShortDesc() string { return `prints message "bye bye"` }
func (t goodbyeCmd) LongDesc() string  { return t.ShortDesc() }
func (t goodbyeCmd) Exec(ctx context.Context, args []string) (context.Context, error) {
	out := ctx.Value("gosh.stdout").(io.Writer)
	fmt.Fprintln(out, "bye bye")
	return ctx, nil
}

// command module
type testCmds struct{}

func (t *testCmds) Init(ctx context.Context) error {
	out := ctx.Value("gosh.stdout").(io.Writer)
	fmt.Fprintln(out, "test module loaded OK")
	return nil
}

func (t *testCmds) Registry() map[string]api.Command {
	return map[string]api.Command{
		"hello":   helloCmd("hello"),
		"goodbye": goodbyeCmd("goodbye"),
	}
}

var Commands testCmds
