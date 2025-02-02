package commands

import (
	"mTLS/ssl"
	"os"
	"path/filepath"
)

type GenerateCaCmd struct {
	Configuration string `arg:"-c,--config" help:"Path to the configuration file"`
}

// GenerateCa - handles the generation of a Certificate Authority.
func GenerateCa(args *GenerateCaCmd) {

	// Validate the arguments
	//
	if args.Configuration == "" {
		panic("Configuration file is required")
	}
	if _, err := os.Stat(args.Configuration); os.IsNotExist(err) {
		panic("Configuration file does not exist")
	}
	if filepath.Ext(args.Configuration) != ".yaml" {
		panic("Configuration file must be in YAML format")
	}

	config, err := ssl.MakeCaConfigurationYaml(args.Configuration)
	if err != nil {
		panic(err)
	}

	println(config.Output.Overwrite)

}
