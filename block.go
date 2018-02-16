package main

import (
    "encoding/json"
)

type Block struct {
    PrevHash string
    Transactions []Transaction
    Height uint64
    Time uint64
    // Nonce string
    // Solution string
}

func (b *Block) Serialize() ([]byte, error){
    return json.Marshal(b)
}

func Deserialize(buf []byte) (*Block, error){
    var blk Block
    err := json.Unmarshal(buf, &blk)
    if err != nil {
        return nil, err
    }
    return &blk, nil
}
