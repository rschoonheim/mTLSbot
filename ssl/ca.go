package ssl

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
)

// X509 - transforms the CaConfiguration into an x509 certificate.
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

// GenerateCa - generates a Certificate Authority.
func (config CaConfiguration) GenerateCa() (*x509.Certificate, []byte, []byte) {
	x509Cert := config.X509()

	caPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, x509Cert, x509Cert, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		panic(err)
	}

	// PEM encode the private key
	//
	caPEM := new(bytes.Buffer)
	pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})

	caPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivateKey),
	})

	// Return the x509 certificate, the private key, and the public key
	//
	return x509Cert, caPrivKeyPEM.Bytes(), caPEM.Bytes()
}

// GenerateKeys - generates a private key and a public key.
func GenerateKeys() (privateKey, publicKey []byte) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	privateKey = x509.MarshalPKCS1PrivateKey(key)
	publicKey, err = x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		panic(err)
	}

	// todo: pem encode the keys

	// Pem encode the private key
	//
	privKeyPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKey,
	}

	// Pem encode the public key
	//
	pubKeyPem := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKey,
	}

	// Encode the private key
	//
	privateKey = pem.EncodeToMemory(privKeyPem)

	// Encode the public key
	//
	publicKey = pem.EncodeToMemory(pubKeyPem)

	return privateKey, publicKey
}
