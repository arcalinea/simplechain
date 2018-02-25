package main

import (
  // "github.com/ipfs/go-datastore"
  // dag "github.com/ipfs/go-ipfs/merkledag"
  // ipld "github.com/ipfs/go-ipld-format"
  // bstore "github.com/ipfs/go-ipfs-blockstore"
)

// func NewBlockchain() {
//   dstore := datastore.NewMapDatastore()
//   blocks := bstore.NewBlockstore(dstore)
//   dagserv := dag.NewDAGService(blocks)
// }

type Blockchain struct {
    Head            *Block
    // ChainDB         ipld.DAGService
    GenesisBlock    *Block
}

func CreateGenesisBlock() (*Block){
    genesisBlock := &Block{ 
    Height: 0,
    Time: 42,
    }
    return genesisBlock
}

func (chain *Blockchain) GetChainTip() (*Block){
    return chain.Head
}

func (chain *Blockchain) ValidateBlock(blk *Block) bool {
    chainTip := chain.GetChainTip()
    if blk.PrevHash != chainTip.GetHash() {
        return false
    }
    return true
}

// Check that prevHash of new block is equal to hash of chainTip
// Transactions validate 
// Height is 1 greater than chainTip
// Time is greater than time of chainTip
func (chain *Blockchain) AddBlock(blk *Block){
    if chain.ValidateBlock(blk) {
        blkCopy := *blk
        chain.Head = &blkCopy
    }
}
