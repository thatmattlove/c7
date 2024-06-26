package cmd

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/lipgloss"
)

func Error(ctx *kong.Context, err error) error {
	r := ErrorStyle.Render(err.Error())
	fmt.Fprintln(ctx.Stderr, r)
	defer ctx.Exit(1)
	return err
}

func KeyValue(ctx *kong.Context, valRenderer lipgloss.Style, key, value string) error {
	k := KeyStyle.Width(len(value) + 3).Render(key)
	v := valRenderer.Render(value)
	c := lipgloss.JoinVertical(lipgloss.Left, k, v)
	f := Container.Render(c)
	fmt.Fprintln(ctx.Stdout, f)
	defer ctx.Exit(0)
	return nil
}
