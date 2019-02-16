package main

import (
	"context"
	"fmt"
	"github.com/vladimirvivien/gosh/api"
	"io"
	"strconv"
	"time"
)

type sleepCmd string

func (s sleepCmd) Name() string {
	return string(s)
}
func (s sleepCmd) Usage() string     { return "Usage: sleep <duration>" }
func (s sleepCmd) ShortDesc() string { return "sleeps for <duration> seconds" }

func (s sleepCmd) LongDesc() string { return s.ShortDesc() }

func (s sleepCmd) Exec(ctx context.Context, args []string) (context.Context, error) {
	if len(args) == 2 {
		duration, err := strconv.Atoi(args[1])
		if err != nil {
			return ctx, err
		}
		time.Sleep(time.Duration(duration) * time.Second)
		return ctx, nil
	}
	out := ctx.Value("gosh.stdout").(io.Writer)
	fmt.Fprintln(out, s.Usage())
	return ctx, nil

}

type sleepCmds struct{}

func (s *sleepCmds) Init(ctx context.Context) error {
	out := ctx.Value("gosh.stdout").(io.Writer)
	fmt.Fprintln(out, "sleep module loaded")
	return nil
}

func (s *sleepCmds) Registry() map[string]api.Command {
	return map[string]api.Command{
		"sleep": sleepCmd("sleep"),
	}
}

var Commands sleepCmds
