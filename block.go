package main

import (
    "encoding/json"
    "crypto/sha256"
)

type Block struct {
    PrevHash [32]byte
    Transactions []Transaction
    Height uint64
    Time uint64
    // Nonce string
    // Solution string
}

func (b *Block) Serialize() ([]byte, error){
    return json.Marshal(b)
}

func DeserializeBlock(buf []byte) (*Block, error){
    var blk Block
    err := json.Unmarshal(buf, &blk)
    if err != nil {
        return nil, err
    }
    return &blk, nil
}

func (b *Block) GetHash() ([32]byte){
    ser, err := b.Serialize() 
    if err != nil {
        panic(err)
    }
    return sha256.Sum256(ser)
}
