package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"

	"github.com/vladimirvivien/gosh/api"
)

// helpCmd represents the `help` command
// which prints out help information about other commands
type helpCmd string

func (h helpCmd) Name() string     { return string(h) }
func (h helpCmd) Usage() string    { return fmt.Sprintf("%s or %s <command-name>", h.Name(), h.Name()) }
func (h helpCmd) LongDesc() string { return "" }
func (h helpCmd) ShortDesc() string {
	return `prints help information for other commands.`
}

func (h helpCmd) Exec(ctx context.Context, args []string) (context.Context, error) {
	if ctx == nil {
		return ctx, errors.New("nil context")
	}

	out := api.GetStdout(ctx)

	cmdsVal := ctx.Value("gosh.commands")
	if cmdsVal == nil {
		return ctx, errors.New("nil context")
	}

	commands, ok := cmdsVal.(map[string]api.Command)
	if !ok {
		return ctx, errors.New("command map type mismatch")
	}

	var cmdNameParam string
	if len(args) > 1 {
		cmdNameParam = args[1]
	}

	// print help for a specified command
	if cmdNameParam != "" {
		cmd, found := commands[cmdNameParam]
		if !found {
			str := fmt.Sprintf("command %s not found", cmdNameParam)
			return ctx, errors.New(str)
		}
		fmt.Fprintf(out, "\n%s\n", cmdNameParam)
		if cmd.Usage() != "" {
			fmt.Fprintf(out, "  Usage: %s\n", cmd.Usage())
		}
		if cmd.ShortDesc() != "" {
			fmt.Fprintf(out, "  %s\n\n", cmd.ShortDesc())
		}
		if cmd.LongDesc() != "" {
			fmt.Fprintf(out, "%s\n\n", cmd.LongDesc())
		}
		return ctx, nil
	}

	fmt.Fprintf(out, "\n%s: %s\n", h.Name(), h.ShortDesc())
	fmt.Fprintln(out, "\nAvailable commands")
	fmt.Fprintln(out, "------------------")
	for cmdName, cmd := range commands {
		fmt.Fprintf(out, "%12s:\t%s\n", cmdName, cmd.ShortDesc())
	}
	fmt.Fprintln(out, "\nUse \"help <command-name>\" for detail about the specified command\n")
	return ctx, nil
}

// exitCmd implements a command to exit the shell
type exitCmd string

func (c exitCmd) Name() string     { return string(c) }
func (c exitCmd) Usage() string    { return "exit" }
func (c exitCmd) LongDesc() string { return "" }
func (c exitCmd) ShortDesc() string {
	return `exits the interactive shell immediately`
}
func (c exitCmd) Exec(ctx context.Context, args []string) (context.Context, error) {
	out := api.GetStdout(ctx)
	fmt.Fprintln(out, "exiting...")
	os.Exit(0)
	return ctx, nil
}

// promptCmd a command that can change the prompt value
type promptCmd string

func (c promptCmd) Name() string     { return string(c) }
func (c promptCmd) Usage() string    { return "prompt <new-prompt>" }
func (c promptCmd) LongDesc() string { return "" }
func (c promptCmd) ShortDesc() string {
	return `sets a new shell prompt`
}
func (c promptCmd) Exec(ctx context.Context, args []string) (context.Context, error) {
	if len(args) < 2 {
		return ctx, errors.New("unable to set prompt, see usage")
	}
	return context.WithValue(ctx, "gosh.prompt", args[1]), nil
}

// sysinfoCmd implements a command that returns system information
type sysinfoCmd string

func (c sysinfoCmd) Name() string     { return string(c) }
func (c sysinfoCmd) Usage() string    { return c.Name() }
func (c sysinfoCmd) LongDesc() string { return "" }
func (c sysinfoCmd) ShortDesc() string {
	return `sets a new shell prompt`
}
func (c sysinfoCmd) Exec(ctx context.Context, args []string) (context.Context, error) {
	out := api.GetStdout(ctx)

	hostname, _ := os.Hostname()
	exe, _ := os.Executable()
	memStats := new(runtime.MemStats)
	runtime.ReadMemStats(memStats)
	info := []struct{ name, value string }{
		{"arc", runtime.GOARCH},
		{"os", runtime.GOOS},
		{"cpus", strconv.Itoa(runtime.NumCPU())},
		{"mem", strconv.FormatUint(memStats.Sys, 10)},
		{"hostname", hostname},
		{"pagesize", strconv.Itoa(os.Getpagesize())},
		{"groupid", strconv.Itoa(os.Getgid())},
		{"userid", strconv.Itoa(os.Geteuid())},
		{"pid", strconv.Itoa(os.Getpid())},
		{"exec", exe},
	}

	fmt.Fprint(out, "\nSystem Info")
	fmt.Fprint(out, "\n-----------")
	for _, k := range info {
		fmt.Fprintf(out, "\n%12s:\t%s", k.name, k.value)
	}
	fmt.Fprintln(out, "\n")
	return ctx, nil
}

// sysCommands represents a collection of commands supported by this
// command module.
type sysCommands struct {
	stdout io.Writer
}

func (c *sysCommands) Init(ctx context.Context) error {
	return nil
}

func (t *sysCommands) Registry() map[string]api.Command {
	return map[string]api.Command{
		"help":   helpCmd("help"),
		"exit":   exitCmd("exit"),
		"prompt": promptCmd("prompt"),
		"sys":    sysinfoCmd("sys"),
	}
}

// plugin entry point
var Commands sysCommands
