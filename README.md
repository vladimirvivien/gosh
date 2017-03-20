# Gosh - A pluggable interactive shell written Go

`Gosh` (or Go shell) is a framework that uses Go's plugin system to create
for building interactive console-based shell programs.  A gosh shell is
comprised of a collection of Go plugins which implement one or more commands.
When `gosh` starts, it searches director `./plugins` for available shared object
files that implement command plugins.

## Getting started
Gosh makes it easy to create a shell programs.  First, download or clone this 
repository.  For a quick start, run the following:

```bash
go run shell/gosh.go
```
This will produce the following output:
```bash

                        888
                        888
                        888
 .d88b.  .d88b. .d8888b 88888b.
d88P"88bd88""88b88K     888 "88b
888  888888  888"Y8888b.888  888
Y88b 888Y88..88P     X88888  888
 "Y88888 "Y88P"  88888P'888  888
     888
Y8b d88P
 "Y88P"

No commands found
```
After the splashscreen is displayed, `gosh` informs you that `no commands found`, as expected.  Next,
exit the `gosh` shell (`Ctrl-C`) and let us compile the example plugins that comes with the source code.

```bash
go build -buildmode=plugin  -o plugins/sys_command.so plugins/syscmd.go
```
The previous command will compile `plugins/syscmd.go` and outputs shared object
`plugins/sys_command.so`, as a Go plugin file.  Verify the shared object file was created:

```
> ls -lh plugins/
total 3.2M
-rw-rw-r-- 1  4.5K Mar 19 18:23 syscmd.go
-rw-rw-r-- 1  3.2M Mar 19 19:14 sys_command.so
-rw-rw-r-- 1  1.4K Mar 19 18:23 testcmd.go
```
Now, when gosh is restarted, it will dynamically load the commands implemented in the shared object file:

```bash
> go run shell/gosh.go
...

Loaded 4 command(s)...
Type help for available commands

gosh>
```

As indicated, typing `help` lists all available commands in the shell:

```bash
gosh> help

help: prints help information for other commands.

Available commands
------------------
      prompt:	sets a new shell prompt
         sys:	sets a new shell prompt
        help:	prints help information for other commands.
        exit:	exits the interactive shell immediately

Use "help <command-name>" for detail about the specified command
```
## A command
A Gosh `Command` is represented by type `api/Command`:
```go
type Command interface {
	Name() string
	Usage() string
	ShortDesc() string
	LongDesc() string
	Exec(context.Context, []string) (context.Context, error)
}
```

The Gosh framework searches for Go plugin files in the `./plugins` directory.  Each package plugin must 
export a variable named `Commands` which is of type  :
```go
type Commands interface {
  ...
	Registry() map[string]Command
}
```
Type `Commands` type returns a list of `Command` via the `Registry()`.  

The following shows example command file [plugins/testcmd.go](./plugins/testcmd.go). It implements
two commands via types `helloCmd` and `goodbyeCmd`. The commands are exported via type `testCmds` using
method `Registry()`:

```
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
```

## License
MIT
