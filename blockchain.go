package main

import (
  "github.com/ipfs/go-datastore"
  dag "github.com/ipfs/go-ipfs/merkledag"
  ipld "github.com/ipfs/go-ipld-format"
  bstore "github.com/ipfs/go-ipfs-blockstore"
)

func NewBlockchain() {
  dstore := datastore.NewMapDatastore()
  blocks := bstore.NewBlockstore(dstore)
  dagserv := dag.NewDAGService(blocks)
}

type Blockchain struct {
    Head            *Block
    ChainDB         ipld.DAGService
    GenesisBlock    *Block
}
