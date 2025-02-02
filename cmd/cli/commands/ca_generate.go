package commands

import (
	"gopkg.in/yaml.v3"
	"mTLS/ssl"
	"os"
	"path/filepath"
)

type GenerateCaCmd struct {
	Configuration string `arg:"-c,--config" help:"Path to the configuration file"`
}

// GenerateCa - handles the generation of a Certificate Authority.
func GenerateCa(args *GenerateCaCmd) {

	// Is the configuration file path empty?
	//
	if args.Configuration == "" {
		panic("Configuration file is required")
	}

	// Does the configuration file exist?
	//
	if _, err := os.Stat(args.Configuration); os.IsNotExist(err) {
		panic("Configuration file does not exist")
	}

	// Is the configuration file YAML?
	//
	if filepath.Ext(args.Configuration) != ".yaml" {
		panic("Configuration file must be in YAML format")
	}

	// Read the yaml
	//
	config, err := os.ReadFile(args.Configuration)
	if err != nil {
		panic(err)
	}

	// Parse the yaml
	//
	var caConfig ssl.CaConfiguration
	err = yaml.Unmarshal(config, &caConfig)
	if err != nil {
		panic(err)
	}

	// Generate x509 certificate
	//
	cert := caConfig.X509()
	println(cert.Subject.Country)
}
