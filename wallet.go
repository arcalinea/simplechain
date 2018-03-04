package main 

import (
    "math/rand"
)

// Returns mock pubkey as random string
func (node *Node) pubkey() string{
    var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
 
    b := make([]rune, 32)
    for i := range b {
        b[i] = letter[rand.Intn(len(letter))]
    }
    return string(b)
}
