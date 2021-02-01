package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
)

const METHOD = "REGISTER"
const COP = "auth"

func main() {
	Nonce := flag.String("n", "", "nonce")
	Realm := flag.String("r", "", "realm")
	Uri := flag.String("uri", "", "URI")
	Username := flag.String("u", "", "Username")
	Password := flag.String("p", "", "Password")
	ClientNonce := flag.String("cn", "", "client nonce")
	NonceCount := flag.String("nc", "", "nonce count")

	flag.Parse()

	HA1sourcestring := fmt.Sprintf("%s:%s:%s", *Username, *Realm, *Password)
	md5.New()
	fmt.Println("A1:", HA1sourcestring)
	HA1b := md5.Sum([]byte(HA1sourcestring))
	HA1Printable := hex.EncodeToString(HA1b[:])
	fmt.Println("HA1:", HA1Printable)
	md5.New()
	HA2source := fmt.Sprintf("%s:%s", METHOD, *Uri)
	fmt.Println("A2:", HA2source)
	HA2b := md5.Sum([]byte(HA2source))
	HA2Printable := hex.EncodeToString(HA2b[:])
	fmt.Println("HA2:", HA2Printable)

	A3string := ""
	ResponseEncodedString := ""
	if len(*ClientNonce) > 0 && len(*NonceCount) > 0 {
		A3string = fmt.Sprintf("%s:%s:%s:%s:%s:%s", HA1Printable, *Nonce, *NonceCount, *ClientNonce, COP, HA2Printable)

	} else {
		md5.New()
		A3string = fmt.Sprintf("%s:%s:%s", HA1Printable, *Nonce, HA2Printable)

	}
	HAResp := md5.Sum([]byte(A3string))
	ResponseEncodedString = hex.EncodeToString(HAResp[:])
	fmt.Println("Response:", ResponseEncodedString)

}
