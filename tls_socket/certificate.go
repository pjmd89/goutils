package tls_socket

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"
)

func (o *TlsSocketServer) get_certificate() (config *tls.Config, err error) {
	cert, err := o.generate_tls_certificate()
	if err != nil {
		return
	}
	config = &tls.Config{Certificates: []tls.Certificate{cert}}
	return
}

func (o *TlsSocketServer) generate_tls_certificate() (outCert tls.Certificate, err error) {
	now := time.Now()
	template := &x509.Certificate{
		SerialNumber: big.NewInt(now.Unix()),
		Subject: pkix.Name{
			CommonName:         o.CommonName,                   //"golang federator controller"
			Country:            []string{o.Country},            //[]string{"VE"}
			Organization:       []string{o.Organization},       //[]string{"https://github.com/allsoftwaretech/go-federator-controller"}
			OrganizationalUnit: []string{o.OrganizationalUnit}, //[]string{"All Software"}
		},
		NotBefore:             now,
		NotAfter:              now.AddDate(10, 0, 0),
		SubjectKeyId:          []byte{113, 117, 105, 99, 107, 115, 101, 114, 118, 101},
		BasicConstraintsValid: true,
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}

	cert, err := x509.CreateCertificate(rand.Reader, template, template, priv.Public(), priv)
	if err != nil {
		return
	}

	outCert.Certificate = append(outCert.Certificate, cert)
	outCert.PrivateKey = priv

	return
}
