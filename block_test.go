package main 

import (
    "fmt"
    "testing"
)

func TestHashBlock(t* testing.T){
    blk := &Block{
        PrevHash: "cat",
        Transactions: nil,
        Height: 100,
        Time: 42,
    }
    hash := blk.GetHash()
    fmt.Println(hash)
    if len(hash) != 32 {
        t.Fatal("Oh no my test failed >:(j")
    }
}
