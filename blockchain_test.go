package main 

import (
    "testing"
    "fmt"
)

var (
    head = &Block {
        Height: 0,
    }
    candidate = &Block {
    PrevHash: head.GetHash(),
    Height: 1,
    }
    
)

func TestValidateBlock(t *testing.T){
    var chain Blockchain
    chain.Head = head 
    if !chain.ValidateBlock(candidate){
        t.Fatal("Candidate block was not valid when expected to be", candidate.Serialize())
    }
}

func TestAddBlock(t *testing.T){
    var chain Blockchain 
    chain.Head = head
    chain.AddBlock(candidate)
    if chain.Head.GetHash() != candidate.GetHash() {
        t.Fatal("Add block failed when it was expected to pass")
    }
}

func TestPointers(t *testing.T) {
    var blk *Block
    takesBar(&blk)
    fmt.Println(blk.Serialize())
}

func takesBar(b **Block)  {
    x := &Block{
        Height: 97,
    }
    *b = x
}

func TestMemory(t *testing.T) {
    buf := make([]byte, 128)
    for i := 0; i < 10; i++ {
        getMessageFromNetwork(&buf)
        fmt.Println(string(buf))
    }
}

func getMessageFromNetwork(buf *[]byte) {
    n := copy(*buf, []byte("cats"))
    *buf = (*buf)[:n]
}
