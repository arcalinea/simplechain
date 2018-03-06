package main

import (
	"crypto/sha256"
	"encoding/json"

	cbor "gx/ipfs/QmRVSCwQtW1rjHCay9NqKXDwbtKTgDcN4iY7PrpSqfKM5D/go-ipld-cbor"
	mh "gx/ipfs/QmZyZDi491cCNTLfAhwcaDii2Kg4pwKRkhqQzURGDvY6ua/go-multihash"
	"gx/ipfs/QmcZfnkapfECQGcLZaf9B79NRg7cRa9EnZh4LSbkCzwNvY/go-cid"
)

type Block struct {
	PrevHash     *cid.Cid
	Transactions []Transaction
	Height       uint64
	Time         uint64
	Nonce        []byte
	Solution     string
}

func (b *Block) Serialize() []byte {
	data, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return data
}

func DeserializeBlock(buf []byte) (*Block, error) {
	var blk Block
	err := json.Unmarshal(buf, &blk)
	if err != nil {
		return nil, err
	}
	return &blk, nil
}

func (b *Block) GetCid() *cid.Cid {
	nd, err := cbor.WrapObject(b, mh.SHA2_256, -1)
	if err != nil {
		panic(err)
	}

	return nd.Cid()
}

func (b *Block) GetHash() [32]byte {
	return sha256.Sum256(b.Serialize())
}
