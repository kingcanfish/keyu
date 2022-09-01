package utils

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

// EmptyArg print empty command hint
func EmptyArg(ctx *cli.Context) {
	fmt.Printf("Empty Content Arg, Command:%v\n", ctx.Command.Name)
}

// ErrArg print empty command hint
func ErrArg(ctx *cli.Context) {
	fmt.Printf("Error Arg, Command:%v, Arg:%v", ctx.Command.Name, ctx.Args().Slice())
}
