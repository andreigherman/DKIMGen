// dkim
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/pem"
	"fmt"
)

func generateDKIM() {
	privatekey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err == nil {
		var publickey *rsa.PublicKey
		publickey = &privatekey.PublicKey
		Priv := x509.MarshalPKCS1PrivateKey(privatekey)
		privBytes := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: Priv,
		})
		privkey := string(privBytes)
		privkeyplain := b64.StdEncoding.EncodeToString([]byte(Priv))
		fmt.Println(Priv)
		fmt.Println()
		fmt.Println(privkeyplain)
		fmt.Println()
		fmt.Println(privkey)
		Pub, err := x509.MarshalPKIXPublicKey(publickey)
		if err == nil {
			pubBytes := pem.EncodeToMemory(&pem.Block{
				Type:  "RSA PUBLIC KEY",
				Bytes: Pub,
			})
			pubkey := string(pubBytes)
			pubkeyplain := b64.StdEncoding.EncodeToString([]byte(Pub))
			fmt.Println(Pub)
			fmt.Println()
			fmt.Println(pubkeyplain)
			fmt.Println()
			fmt.Println(pubkey)
		}
	}
}

func main() {
	generateDKIM()
}
