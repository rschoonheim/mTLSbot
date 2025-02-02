package ssl

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
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

func (config CaConfiguration) X509() *x509.Certificate {
	return &x509.Certificate{
		SerialNumber: config.SerialNumber,
		Subject: pkix.Name{
			Country:            []string{config.Subject.Country},
			Organization:       []string{config.Subject.Organization},
			OrganizationalUnit: []string{config.Subject.OrganizationalUnit},
			Locality:           []string{config.Subject.Locality},
			Province:           []string{config.Subject.Province},
			StreetAddress:      []string{config.Subject.StreetAddress},
			PostalCode:         []string{config.Subject.PostalCode},
		},
		NotBefore:             config.NotBefore,
		NotAfter:              config.NotAfter,
		IsCA:                  config.IsCA,
		ExtKeyUsage:           config.ExtKeyUsage,
		KeyUsage:              config.KeyUsage,
		BasicConstraintsValid: config.BasicConstraintsValid,
	}
}
