/*
** keypair.go
**
** cs4501 Fall 2015
** Problem Set 1
*/

// Every file in go is a part of some package. Since we want to create a program
// from this file we use 'package main'
package main

import (
  "regexp"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
)



func main() {

	// In order to recieve coins we must generate a public/private key pair.
	pub, priv := generateKeyPair()

	// To use this address we must store our private key somewhere. Everything
	// else can be recovered from the private key.
	fmt.Printf("This is a private key in hex:\t[%s]\n",
		hex.EncodeToString(priv.Serialize()))

	fmt.Printf("This is a public key in hex:\t[%s]\n",
		hex.EncodeToString(pub.SerializeCompressed()))

	addr := generateAddr(pub)

	// Output the bitcoin address derived from the public key
	fmt.Printf("This is the associated Bitcoin address:\t[%s]\n", addr.String())
}
