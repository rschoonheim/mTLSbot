package ssl

import (
	"crypto/x509"
	"errors"
	"gopkg.in/yaml.v3"
	"math/big"
	"os"
	"reflect"
	"strings"
	"time"
)

type CaConfiguration struct {
	SerialNumber *big.Int `yaml:"serial-number"`
	Subject      struct {
		Country            string `yaml:"country"`
		Organization       string `yaml:"organization"`
		OrganizationalUnit string `yaml:"organizational-unit"`
		Locality           string `yaml:"locality"`
		Province           string `yaml:"province"`
		StreetAddress      string `yaml:"street-address"`
		PostalCode         string `yaml:"postal-code"`
	} `yaml:"subject"`
	NotBefore             time.Time          `yaml:"not-before"`
	NotAfter              time.Time          `yaml:"not-after"`
	IsCA                  bool               `yaml:"is-ca"`
	ExtKeyUsage           []x509.ExtKeyUsage `yaml:"ext-key-usage"`
	KeyUsage              x509.KeyUsage      `yaml:"key-usage"`
	BasicConstraintsValid bool               `yaml:"basic-constraints-valid"`
}

type CaConfigurationYaml struct {
	Output struct {
		Directory  string `yaml:"directory"`
		Overwrite  bool   `yaml:"overwrite,omitempty"`
		CaCertName string `yaml:"ca-cert-name,omitempty"`
		CaKeyName  string `yaml:"ca-key-name,omitempty"`
	}
	X509 CaConfiguration `yaml:"x509"`
}

// setDefaultValue - sets the default value for a field in the CaConfigurationYaml struct.
func setDefaultValue(caConfig *CaConfigurationYaml, key string, value interface{}) error {
	// Split the key into its parts
	//
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return errors.New("INVALID_KEY")
	}

	// Traverse the keys to find the target
	//
	target := reflect.ValueOf(caConfig).Elem()
	for _, k := range keys[:len(keys)-1] {
		target = target.FieldByName(k)
		if !target.IsValid() {
			return errors.New("INVALID_KEY")
		}
	}

	// Set the default value
	//
	field := target.FieldByName(keys[len(keys)-1])
	if field.IsZero() {
		field.Set(reflect.ValueOf(value))
	}

	return nil
}

// MakeCaConfigurationYaml - transforms the CA configuration YAML into a CaConfigurationYaml struct.
func MakeCaConfigurationYaml(path string) (*CaConfigurationYaml, error) {

	// Start with reading the configuration file and parsing it
	// into a CaConfiguration struct.
	//
	config, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var caConfig CaConfigurationYaml
	err = yaml.Unmarshal(config, &caConfig)
	if err != nil {
		return nil, err
	}

	// Handle setting default values for the configuration
	// when they are not provided.
	//
	defaults := map[string]interface{}{
		"Output.CaCertName": "ca.crt",
		"Output.CaKeyName":  "ca.key",
		"Output.Overwrite":  false,
	}

	for key, value := range defaults {
		if err := setDefaultValue(&caConfig, key, value); err != nil {
			return nil, err
		}
	}

	// Validate `Output.Directory`
	//
	if caConfig.Output.Directory == "" || caConfig.Output.CaCertName == "" || caConfig.Output.CaKeyName == "" {
		return nil, errors.New("OUTPUT_DIRECTORY_INVALID")
	}

	// When `Output.Overwrite` is set to false, check if the `Output.CaCertName` and `Output.CaKeyName` files exist.
	//
	if !caConfig.Output.Overwrite {
		if _, err := os.Stat(caConfig.Output.Directory + "/" + caConfig.Output.CaCertName); !os.IsNotExist(err) {
			return nil, errors.New("CA_CERT_ALREADY_EXISTS")
		}
		if _, err := os.Stat(caConfig.Output.Directory + "/" + caConfig.Output.CaKeyName); !os.IsNotExist(err) {
			return nil, errors.New("CA_KEY_ALREADY_EXISTS")
		}
	}

	return &caConfig, nil
}
