package main

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"
)

var (
	testPluginsDir = "../plugins"
)

func TestShellNew(t *testing.T) {
	shell := New()
	if shell.pluginsDir != pluginsDir {
		t.Error("pluginsDir not set")
	}
}

func TestShellInit(t *testing.T) {
	shell := New()
	shell.pluginsDir = testPluginsDir
	ctx := context.WithValue(context.TODO(), "gosh.stdout", os.Stdout)
	if err := shell.Init(ctx); err != nil {
		t.Fatal(err)
	}
	if len(shell.commands) <= 0 {
		t.Error("failed to load plugins from", testPluginsDir)
	}
	if _, ok := shell.commands["hello"]; !ok {
		t.Error("missing 'hello' command from test module")
	}
	if _, ok := shell.commands["goodbye"]; !ok {
		t.Error("missing 'goodbye' command from test module")
	}

}

func TestShellHandle(t *testing.T) {
	shell := New()
	shell.pluginsDir = testPluginsDir

	ctx := context.WithValue(context.TODO(), "gosh.stdout", os.Stdout)
	if err := shell.Init(ctx); err != nil {
		t.Fatal(err)
	}

	helloOut := bytes.NewBufferString("")
	shell.ctx = context.WithValue(context.TODO(), "gosh.stdout", helloOut)
	if err := shell.handle("testhello"); err == nil {
		t.Error("this test should have failed with command not found")
	}
	if err := shell.handle("hello"); err != nil {
		t.Error(err)
	}
	printedOut := strings.TrimSpace(helloOut.String())
	if printedOut != "hello there" {
		t.Error("did not get expected output from testcmd")
	}

	byeOut := bytes.NewBufferString("")
	shell.ctx = context.WithValue(context.TODO(), "gosh.stdout", byeOut)
	if err := shell.handle("goodbye"); err != nil {
		t.Error(err)
	}
	printedOut = strings.TrimSpace(byeOut.String())
	if printedOut != "bye bye" {
		t.Error("did not get expected output from testcmd")
	}

}
