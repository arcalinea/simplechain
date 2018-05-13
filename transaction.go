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

type TxMerkleTree struct {
    RootNode *MerkleNode
}

type MerkleNode struct {
    Left *MerkleNode
    Right *MerkleNode 
    Data []byte
}

func NewMerkleNode(left *MerkleNode, right *MerkleNode, tx *Transaction) *MerkleNode {
    mnode := MerkleNode{}
    
    if left == nil && right == nil {
        hash := sha256.Sum256(tx.Serialize())
        mnode.Data = hash[:]
    } else {
        prevHashes := append(left.Data, right.Data...)
        hash := sha256.Sum256(prevHashes)
        mnode.Data = hash[:]
    }
    mnode.Left = left 
    mnode.Right = right 
    
    return &mnode
}

func BuildTxMerkleTree(txs []*Transaction) *TxMerkleTree {
    var nodes []MerkleNode
    
    node := NewMerkleNode(nil, nil, txs[0])
    nodes = append(nodes, *node)
    
    for i := 0; i < len(txs)/2; i++ {
        var newLevel []MerkleNode
        for j := 0; j < len(nodes); j += 2 {
            node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
            newLevel = append(newLevel, *node)
        }
        nodes = newLevel
    }
    mtree := TxMerkleTree{&nodes[0]}
    return &mtree
}

////

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
