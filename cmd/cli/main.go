package main

import (
	"github.com/alexflint/go-arg"
	"mTLS/cmd/cli/commands"
)

type args struct {
	GenerateCa *commands.GenerateCaCmd `arg:"subcommand:generate-ca"`
}

func (args) Version() string {
	return "Version: 1.0.0"
}

func (args) Description() string {
	return "mTLS bot provides you with utilities to manage mTLS certificates"
}

func main() {
	var args args
	arg.MustParse(&args)

	switch {
	case args.GenerateCa != nil:
		commands.GenerateCa(args.GenerateCa)
	}

}
