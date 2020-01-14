package main

// this is on the back burner, not needed at the moment

import (
	"context"
	"fmt"

	"github.com/vladimirvivien/gosh/api"
)

type splashCmd string
type splashCmds struct{}

func (t *splashCmds) Init(ctx context.Context) error {
	// to set your splash, modify the text in the println statement below, multiline is supported
	fmt.Println(`		
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
 	
 `)

	return nil
}

func (t *splashCmds) Registry() map[string]api.Command {
	return map[string]api.Command{}
}

var Commands splashCmds
