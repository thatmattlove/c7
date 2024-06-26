package cmd

import (
	"github.com/alecthomas/kong"
)

type CLI struct {
	Encode Encode `cmd:"encode" aliases:"e" help:"${encode_help}"`
	Decode Decode `cmd:"decode" aliases:"d" help:"${decode_help}"`
}

func Run() {
	cli := CLI{Encode{}, Decode{}}
	name := TitleStyle.Render("c7")
	desc := DescriptionStyle.Render("Quickly Encode/Decode a Cisco Type 7 Password")

	ctx := kong.Parse(&cli,
		kong.Name(name),
		kong.Description(desc),
		kong.Vars{
			"encode_help": DescriptionStyle.Render("Encode a value to a Cisco type 7 hash"),
			"decode_help": DescriptionStyle.Render("Decode a Cisco type 7 hash back to plain text"),
		},
	)
	err := ctx.Run(&kong.Context{})
	if err != nil {
		ErrorStyle.Render(err.Error())
		ctx.Exit(1)
	}
}
