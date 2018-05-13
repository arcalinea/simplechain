package main 

import (
    "fmt"
    "testing"
)

func TestToAddress(t *testing.T){
    new := newKey()
    addr := ToAddress(new.PublicKey)
    fmt.Println("Address:", addr)
}
