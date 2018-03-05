package main

import (
    "encoding/json"
    "encoding/hex"
    "crypto/sha256"
)

type Transaction struct {
    Sender string
    Receiver string
    Amount uint64
    Memo string
}

func (tx *Transaction) Serialize() ([]byte){
    data, err := json.Marshal(tx)
    if err != nil {
        panic(err)
    }
    return data
}

func DeserializeTx(buf []byte) (*Transaction, error){
    var tx Transaction
    err := json.Unmarshal(buf, &tx)
    if err != nil {
        return nil, err
    }
    return &tx, nil
}

func (tx *Transaction) GetTxid() ([32]byte) {
    ser := tx.Serialize() 
    return sha256.Sum256(ser)
}

func (tx *Transaction) GetTxidString() string {
    txid := tx.GetTxid()
    return hex.EncodeToString(txid[:])
}
