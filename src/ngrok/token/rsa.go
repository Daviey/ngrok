package token

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

var priv *rsa.PrivateKey
var pub *rsa.PublicKey

func readKeys() ([]byte, []byte) {
	privateKey, err := ioutil.ReadFile("assets/token/private.pem")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading private file, does it exist? (./assets/token/private.pem)\n")
		os.Exit(1)
	}

	publicKey, err := ioutil.ReadFile("assets/token/public.pem")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading publickey file, does it exist? (./assets/token/public.pem)\n")
		os.Exit(1)
	}

	return privateKey, publicKey

}

func init() {
	// TBD handle error
	privateKey, publicKey := readKeys()

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
