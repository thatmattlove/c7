package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

type argT struct {
	Help   bool `cli:"h,help" usage:"show help"`
	Encode bool `cli:"e,encode" usage:"Encode instead of decode"`
}

// AutoHelp implements cli.AutoHelper interface
// NOTE: cli.Helper is a predefined type which implements cli.AutoHelper
func (argv *argT) AutoHelp() bool {
	return argv.Help
}

func main() {
	args := os.Args[1:]

	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		if len(args) == 0 {
			return fmt.Errorf("%s", ctx.Color().Red(ctx.Color().Bold("No argument provided")))
		}
		if len(args) > 1 {
			return fmt.Errorf("%s", ctx.Color().Red(ctx.Color().Bold("Too many arguments provided")))
		}
		if argv.Encode {
			return fmt.Errorf("%s", ctx.Color().Red(ctx.Color().Bold("Not implemented yet")))
		}
		if len(args) == 1 {
			decoded := decode7(args[0])
			ctx.String("\n%s: %s\n", ctx.Color().White("Encoded"), ctx.Color().Red(ctx.Color().Bold(args[0])))
			ctx.String("%s: %s\n", ctx.Color().White("Decoded"), ctx.Color().Green(ctx.Color().Bold(decoded)))
			return nil
		}
		return nil
	}, "\nDecode (or Encode) Cisco Type 7 Passwords"))
}
