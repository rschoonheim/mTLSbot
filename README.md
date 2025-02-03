# mTLS bot - simple mTLS certificate management

**Work in progress**

mTLS bot enables you to easily manage mTLS certificates for your applications. It is designed to be simple and easy to
use. The tool is written in Go and uses the `crypto/x509` package to generate certificates.

## Using mTLS bot


### 1. Creating the root Certificate Authority

### 1.2 Creating an intermediate Certificate Authority (optional)

### 2. Creating server certificates

### 3. Creating client certificates

## Configurations

### Certificate Configuration File

The CA configuration file serves two purposes:

1. It specifies where the CA certificate and key are stored.
2. It specifies the CA certificate configuration.

**Configuration Example**

```yaml
output:
  directory: ".example/output"
  overwrite: false
  ca-cert-name: "ca.crt"
  ca-key-name: "ca.key"
x509:
  serial-number: 1
  subject:
    organization: "Example Inc."
    organization-unit: "Example Unit"
    country: "NL"
    province: "Utrecht"
    locality: "Utrecht"
    streetAddress: "1234 Example St."
    postalCode: "1234AB"
  not-before: "2025-01-02T15:04:05Z"
  not-after: "2035-01-02T15:04:05Z"
  is-ca: true
  ext-key-usage:
    - 1
    - 0
  key-usage: 5
  basic-constraints-validity: true
```

### Server Configuration File
