package main

import (
	"context"
	"time"

	bserv "github.com/ipfs/go-ipfs/blockservice"
	"github.com/ipfs/go-ipfs/exchange/bitswap"
	"github.com/ipfs/go-ipfs/exchange/bitswap/network"
	"gx/ipfs/QmPpegoMqhAEqjncrzArm7KVWAkCm78rqL2DPuNjhPrshg/go-datastore"
	bstore "gx/ipfs/QmTVDM4LCSUMFNQzbDLL9zQwp8usE6QHymFdh3h8vL9v6b/go-ipfs-blockstore"
	nonerouting "gx/ipfs/QmZRcGYvxdauCd7hHnMYLYqcZRaDjv24c7eUNyJojAcdBb/go-ipfs-routing/none"
	mh "gx/ipfs/QmZyZDi491cCNTLfAhwcaDii2Kg4pwKRkhqQzURGDvY6ua/go-multihash"

	cbor "gx/ipfs/QmRVSCwQtW1rjHCay9NqKXDwbtKTgDcN4iY7PrpSqfKM5D/go-ipld-cbor"
	cid "gx/ipfs/QmcZfnkapfECQGcLZaf9B79NRg7cRa9EnZh4LSbkCzwNvY/go-cid"

	host "gx/ipfs/QmNmJZL7FQySMtE2BQuLMuZg2EB2CLEunJJUSVSc9YnnbV/go-libp2p-host"
)

type Blockchain struct {
	Head         *Block
	ChainDB      bserv.BlockService
	GenesisBlock *Block
	Blockstore   bstore.Blockstore
}

func init() {
	cbor.RegisterCborType(Block{})
	cbor.RegisterCborType(Transaction{})
}

func NewBlockchain(h host.Host) *Blockchain {
	// base backing datastore, currently just in memory, but can be swapped out
	// easily for leveldb or other
	dstore := datastore.NewMapDatastore()

	// wrap the datastore in a 'content addressed blocks' layer
	blocks := bstore.NewBlockstore(dstore)

	// now heres where it gets a bit weird. Its currently rather annoying to set up a bitswap instance.
	// Bitswap wants a datastore, and a 'network'. Bitswaps network instance
	// wants a libp2p node and a 'content routing' instance. We don't care
	// about content routing right now, so we want to give it a dummy one.
	// TODO: make bitswap easier to construct
	nr, _ := nonerouting.ConstructNilRouting(nil, nil, nil)
	bsnet := network.NewFromIpfsHost(h, nr)

	bswap := bitswap.New(context.Background(), h.ID(), bsnet, blocks, true)

	// Bitswap only fetches blocks from other nodes, to fetch blocks from
	// either the local cache, or a remote node, we can wrap it in a
	// 'blockservice'
	bservice := bserv.New(blocks, bswap)

	genesis := CreateGenesisBlock()
	return &Blockchain{
		GenesisBlock: genesis,
		Head:         genesis,

		ChainDB:    bservice,
		Blockstore: blocks,
	}
}

func LoadBlock(bs bserv.BlockService, c *cid.Cid) (*Block, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	data, err := bs.GetBlock(ctx, c)
	if err != nil {
		return nil, err
	}

	var out Block
	if err := cbor.DecodeInto(data.RawData(), &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func PutBlock(bs bserv.BlockService, blk *Block) (*cid.Cid, error) {
	nd, err := cbor.WrapObject(blk, mh.BLAKE2B_MIN+31, 32)
	if err != nil {
		return nil, err
	}

	if err := bs.AddBlock(nd); err != nil {
		return nil, err
	}

	return nd.Cid(), nil
}


//////

func CreateGenesisBlock() *Block {
	genesisBlock := &Block{
		Height: 0,
		Time:   42,
	}
	return genesisBlock
}

func (chain *Blockchain) GetChainTip() *Block {
	return chain.Head
}

func validateTransactions(txs []Transaction) bool {
    // TODO: Tx validation logic goes here
    return true
}

// Check that prevHash of new block is equal to hash of chainTip
// Transactions validate
// Height is 1 greater than chainTip
// Time is greater than time of chainTip
func (chain *Blockchain) ValidateBlock(blk *Block) bool {
	chainTip := chain.GetChainTip()
	if blk.PrevHash != chainTip.GetHash() {
		return false
	}
    if !validateTransactions(blk.Transactions){
        return false
    }
    if blk.Height != chainTip.Height + 1 {
        return false
    }
    if blk.Time < chainTip.Time {
        return false
    }
	return true
}

func (chain *Blockchain) AddBlock(blk *Block) {
	if chain.ValidateBlock(blk) {
		blkCopy := *blk
		chain.Head = &blkCopy
	}
}
