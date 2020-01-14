package main

import (
	"context"
	"fmt"
	"io"

	"github.com/vladimirvivien/gosh/api"
)

// Make sure to replace `exampleCmd` with
// what ever you want, but make sure it
// does not clash with existing plugins,

// # MULTIPLE COMMANDS CAN BE DECLARED IN
// # THE SAME PLUGIN, see testcmd.go for
// # an example

type exampleCmd string

// for plugins that dont contain a command to
// load, all lines marked with //OP are optional
// if your plugin adds a command, please fille these out
func (t exampleCmd) Name() string      { return string(t) }                //OP
func (t exampleCmd) Usage() string     { return `example` }                //OP
func (t exampleCmd) ShortDesc() string { return `description of example` } //OP
func (t exampleCmd) LongDesc() string  { return t.ShortDesc() }            //OP
func (t exampleCmd) Exec(ctx context.Context, args []string) (context.Context, error) {

	// Put your custom programming here, make sure to
	// accept and parse the args variable if you need
	// to!

	return ctx, nil
}

type exampleCmds struct{}

func (t *exampleCmds) Init(ctx context.Context) error {
	out := ctx.Value("gosh.stdout").(io.Writer)

	// If you want something to happen when the module
	// loads, put it here,

	fmt.Fprintln(out, "example module loaded OK")
	return nil
}

func (t *exampleCmds) Registry() map[string]api.Command {
	return map[string]api.Command{
		"example": exampleCmd("example"), //OP
	}
}

var Commands exampleCmds

// If your plugin needs extra functions, declare
// them down here to call upon, or import their
// library.
