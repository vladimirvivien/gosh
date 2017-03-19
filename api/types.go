package api

import "context"

// Module a plugin that can be initialized
type Module interface {
	Init(context.Context) error
}

// Command represents an executable a command
type Command interface {
	Name() string
	Usage() string
	ShortDesc() string
	LongDesc() string
	Exec(context.Context, []string) (context.Context, error)
}

// Commands a plugin that contains one or more command
type Commands interface {
	Module
	Registry() map[string]Command
}
