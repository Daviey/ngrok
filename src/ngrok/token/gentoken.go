package token

import (
	"crypto"
	"encoding/base64"
	"flag"
	"fmt"
)

var userdata string

func init() {
}

func Main() {
	flag.StringVar(&userdata, "u", "", "")
	flag.Parse()
	signBytes, err := RsaSign([]byte(userdata), crypto.MD5)
	if err != nil {
		panic(err)
	}
	sign := base64.StdEncoding.EncodeToString(signBytes)
	fmt.Println("new token:" + userdata + "#" + sign)
}
