# Gosh - A pluggable interactive shell written Go

`Gosh` (or Go shell) is a framework that uses Go's plugin system to create
for building interactive console-based shell programs.  A gosh shell is
comprised of a collection of Go plugins which implement one or more commands.
When `gosh` starts, it searches director `./plugins` for available shared object
files that implement command plugins.

## Getting started
Gosh makes it easy to create a shell programs.  First, download or clone this 
repository.  For a quick start, run the following:

```
go run shell/gosh.go
```
This will produce the following output:
```

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

```
go build -buildmode=plugin  -o plugins/sys_command.so plugins/syscmd.go
```
The previous command will compile `plugins/syscmd.go` and outputs shared object
`plugins/sys_command.so`, as a Go plugin file.  Verify the shared object file was created:

```
> ls -l plugins/
total 3256
-rw-rw-r-- 1     4599 Mar 19 18:23 syscmd.go
-rw-rw-r-- 1  3321072 Mar 19 19:14 sys_command.so
-rw-rw-r-- 1     1401 Mar 19 18:23 testcmd.go
```
Now, when gosh is restarted, it will dynamically load the commands implemented in the shared object file:

```
> go run shell/gosh.go

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



Loaded 4 command(s)...
Type help for available commands

gosh>
```

As indicated, typing `help` lists all available commands in the shell:

```
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
