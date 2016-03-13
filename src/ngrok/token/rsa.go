package token 

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIHyAgEAAjEA2P7zpM6LdPFLX+dCIqFqE/TerJiV5YrUdxEJs5wC/Caaz56wdhXc
kAmPylh2ouqXAgMBAAECMQC2ATnonNk5r/2xpe3B0DOCI5v1llH74vGAbvRKlhZa
Ut1FV3ly27kxgbxRq6w4vmkCGQD1TLyN1yrjKQh4aCI/MsvyYo50z5yYBwUCGQDi
diVgn93RwhGhJsq/K0A01XFSvYUp5esCGAizyTiq+n2GliNXZmehkLSvpGgmeWK3
VQIYHk7Xk8XxhvglKd8qNUmRj0CdqQgqQUA/Ahgf9xvUrF4hBWYhAOW16L1yuxWR
y1STgdU=
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MEwwDQYJKoZIhvcNAQEBBQADOwAwOAIxANj+86TOi3TxS1/nQiKhahP03qyYleWK
1HcRCbOcAvwmms+esHYV3JAJj8pYdqLqlwIDAQAB
-----END PUBLIC KEY-----
`)

var priv *rsa.PrivateKey
var pub  *rsa.PublicKey
func init() {
	// TBD handle error
        privblock, _ := pem.Decode(privateKey)
        priv, _ = x509.ParsePKCS1PrivateKey(privblock.Bytes)

	pubblock, _ := pem.Decode(publicKey)
        pubInterface, _ := x509.ParsePKIXPublicKey(pubblock.Bytes)
        pub = pubInterface.(*rsa.PublicKey)	
}
func RsaEncrypt(data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func RsaSign(data []byte, hash crypto.Hash) ([]byte, error) {
	h := hash.New()
	h.Write(data)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, priv, hash, hashed)
}

func RsaVerify(data []byte, sign []byte, hash crypto.Hash) error {
	h := hash.New()
	h.Write(data)
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(pub, hash, hashed, sign)
}

