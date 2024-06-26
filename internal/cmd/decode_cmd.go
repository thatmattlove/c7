package cmd

import (
	"github.com/alecthomas/kong"
	"github.com/thatmattlove/c7/internal/encoding"
)

type Decode struct {
	Value string `arg:"" name:"value" help:"Value to decode"`
}

func (c *Decode) Run(ctx *kong.Context) error {
	result, err := encoding.Decode(c.Value)
	if err != nil {
		return Error(ctx, err)
	}
	return KeyValue(ctx, ValStyle1, "Decoded", result)
}
