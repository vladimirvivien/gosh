package main

// this is on the back burner, not needed at the moment

import (
	"context"
	"fmt"

	"github.com/Merith-TK/gosh/api"
)

type splashCmd string
type splashCmds struct{}

func (t *splashCmds) Init(ctx context.Context) error {
	// to set your splash, modify the text in the println statement below, multiline is supported
	fmt.Println(`-------------------------`)

	return nil
}

func (t *splashCmds) Registry() map[string]api.Command {
	return map[string]api.Command{}
}

var Commands splashCmds
