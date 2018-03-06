package main

import (
	"context"
	"fmt"
	"os"
	"time"

	libp2p "gx/ipfs/QmNh1kGFFdsPu79KNSaL4NUKUPb4Eiz4KHdMtFY6664RDp/go-libp2p"
	host "gx/ipfs/QmNmJZL7FQySMtE2BQuLMuZg2EB2CLEunJJUSVSc9YnnbV/go-libp2p-host"
	ipfsaddr "gx/ipfs/QmQViVWBHbU6HmYjXcdNq7tVASCNgdg64ZGcauuDkLCivW/go-ipfs-addr"
	floodsub "gx/ipfs/QmSFihvoND3eDaAYRCeLgLPt62yCPgMZs1NSZmKFEtJQQw/go-libp2p-floodsub"
	peerstore "gx/ipfs/QmXauCuJzmzapetmC6W4TuDJLL1yFFrVzSHoWv8YdbmnxH/go-libp2p-peerstore"

	"./types"
)

type Node struct {
	p2pNode    host.Host
	mempool    *Mempool
	blockchain *Blockchain
	pubsub     *floodsub.PubSub
}

func CreateNewNode(ctx context.Context) *Node {
	var node Node

	newNode, err := libp2p.New(ctx, libp2p.Defaults)
	if err != nil {
		panic(err)
	}

	pubsub, err := floodsub.NewFloodSub(ctx, newNode)
	if err != nil {
		panic(err)
	}

	for i, addr := range newNode.Addrs() {
		fmt.Printf("%d: %s/ipfs/%s\n", i, addr, newNode.ID().Pretty())
	}

	if len(os.Args) > 1 {
		addrstr := os.Args[1]
		addr, err := ipfsaddr.ParseString(addrstr)
		if err != nil {
			panic(err)
		}
		fmt.Println("Parse Address:", addr)

		pinfo, _ := peerstore.InfoFromP2pAddr(addr.Multiaddr())

		if err := newNode.Connect(ctx, *pinfo); err != nil {
			fmt.Println("bootstrapping a peer failed", err)
		}
	}

	blockchain := NewBlockchain(newNode)

	node.p2pNode = newNode
	node.mempool = NewMempool()
	node.pubsub = pubsub
	node.blockchain = blockchain

	node.ListenBlocks(ctx)
	node.ListenTransactions(ctx)

	return &node

}

func (node *Node) ListenBlocks(ctx context.Context) {
	sub, err := node.pubsub.Subscribe("blocks")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			msg, err := sub.Next(ctx)
			if err != nil {
				panic(err)
			}
			blk, err := DeserializeBlock(msg.GetData())
			if err != nil {
				panic(err)
			}
			// fmt.Println("Block received over network:", string(blk.Serialize()))
			fmt.Println("Block received over network, blockhash", blk.GetCid())
			cid := node.blockchain.AddBlock(blk)
			if cid != nil {
				fmt.Println("Block added, cid:", cid)
				node.mempool.removeTxs(blk.Transactions)
			}
		}
	}()
}

func (node *Node) ListenTransactions(ctx context.Context) {
	sub, err := node.pubsub.Subscribe("transactions")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			msg, err := sub.Next(ctx)
			if err != nil {
				panic(err)
			}
			tx, err := DeserializeTx(msg.GetData())
			if err != nil {
				panic(err)
			}
			node.mempool.AddTx(tx)
			fmt.Println("Tx received over network, added to mempool:", tx)
		}
	}()
}

func (node *Node) CreateNewBlock() *Block {
	var blk Block
	blk.PrevHash = node.blockchain.Head.GetCid()
	blk.Transactions = node.mempool.SelectTransactions()
	blk.Height = node.blockchain.Head.Height + 1
	blk.Time = uint64(time.Now().Unix())
	return &blk
}

func (node *Node) BroadcastBlock(block *Block) {
	data := block.Serialize()
	node.pubsub.Publish("blocks", data)
}

func (node *Node) SendTransaction(tx *Transaction) *types.SendTxResponse {
	var res types.SendTxResponse
	txid := tx.GetTxid()
	node.mempool.transactions[txid] = *tx
	data := tx.Serialize()
	node.pubsub.Publish("transactions", data)
	res.Txid = tx.GetTxidString()
	return &res
}

func (node *Node) GetInfo() *types.GetInfoResponse {
	var res types.GetInfoResponse
	res.BlockHeight = node.blockchain.Head.Height
	return &res
}
