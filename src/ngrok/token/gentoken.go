package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
)

var userdata string

func init() {
	flag.StringVar(&userdata, "u", "", "")
	flag.Parse()
}

func main() {
	var data []byte
	var err error
	data, err = RsaEncrypt([]byte(userdata))
	if err != nil {
		panic(err)
	}
	fmt.Println("new token:" + base64.StdEncoding.EncodeToString(data) + userdata)
}

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIGqAgEAAiEAtvAot8pIjNpjhe2QHvBrE+LrGO2Tyi0mO9J4d1+lNOUCAwEAAQIg
PgnJtTJLUkEJeTSsKGHd3aii3iTQXcdI/dG3TM6RcAECEQDjijGUqATgYxfOv9/8
2LwBAhEAzdHSmyiJKlvmDE0PP+UI5QIRANSu9Y0oR/adaUF9pBYpWAECEFzZZDYk
wkQ+svgnp1J+w1UCECFEYqakMM9jZNTslmvxHac=
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MDwwDQYJKoZIhvcNAQEBBQADKwAwKAIhALbwKLfKSIzaY4XtkB7waxPi6xjtk8ot
JjvSeHdfpTTlAgMBAAE=
-----END PUBLIC KEY-----
`)

func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
