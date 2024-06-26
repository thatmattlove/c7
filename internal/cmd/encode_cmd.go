package cmd

import (
	"github.com/alecthomas/kong"
	"github.com/thatmattlove/c7/internal/encoding"
)

type Encode struct {
	Value string `arg:"" name:"value" help:"Value to encode"`
}

func (c *Encode) Run(ctx *kong.Context) error {
	result, err := encoding.Encode(c.Value)
	if err != nil {
		return Error(ctx, err)
	}
	return KeyValue(ctx, ValStyle2, "Encoded", result)
}
