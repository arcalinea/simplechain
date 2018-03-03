package main 

import (
    // "crypto/ecdsa"
    "math/rand"
)


func (node *Node) pubkey() string{
    var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
 
    b := make([]rune, 32)
    for i := range b {
        b[i] = letter[rand.Intn(len(letter))]
    }
    return string(b)
}
