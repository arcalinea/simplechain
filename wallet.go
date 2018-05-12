package main 

import (
    "crypto/rand"
    "crypto/ecdsa"
    "crypto/elliptic"
)

func newKey() *ecdsa.PrivateKey {
    key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        panic(err)
    }
	return key
}
