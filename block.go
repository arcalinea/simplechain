package main

import (
    "encoding/json"
    "crypto/sha256"
    "encoding/base64"
)

type Block struct {
    PrevHash [32]byte
    Transactions []Transaction
    Height uint64
    Time uint64
    Nonce []byte
    Solution string
}

func (b *Block) Serialize() ([]byte){
    data, err := json.Marshal(b)
    if err != nil {
        panic(err)
    }
    return data
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
    ser := b.Serialize() 
    return sha256.Sum256(ser)
}

func (b *Block) GetHashString() string {
    hash := b.GetHash()
    hashstr := hash[:]
    return base64.StdEncoding.EncodeToString(hashstr)
}
