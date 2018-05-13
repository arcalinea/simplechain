package main 

import (
    "crypto/rand"
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/sha256"
    "encoding/base64"
)


func newKey() *ecdsa.PrivateKey {
    key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        panic(err)
    }
	return key
}

func ToAddress(pubkey ecdsa.PublicKey) string {
    pubBytes := elliptic.Marshal(elliptic.P256(), pubkey.X, pubkey.Y)
    addr := sha256.Sum256(pubBytes[1:])
    addrString := base64.StdEncoding.EncodeToString(addr[:])[12:]
    return addrString
}
