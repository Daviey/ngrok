package token

import (
	"crypto"
	"encoding/base64"
	"flag"
	"fmt"
)

var userdata string

func init() {
	flag.StringVar(&userdata, "u", "", "")
	flag.Parse()
}

func Main() {
	signBytes, err := RsaSign([]byte(userdata), crypto.MD5)
	if err != nil {
		panic(err)
	}
	sign := base64.StdEncoding.EncodeToString(signBytes)
	fmt.Println("new token:" + userdata + "#" + sign)
}
