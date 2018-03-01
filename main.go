package main

import (
	"context"
)

func main() {
	ctx := context.Background()
	
	node := CreateNewNode(ctx)
	
	var blk Block
	blk.PrevHash = node.blockchain.Head.GetHash()
	blk.Transactions = []Transaction{
		{Sender: "arc", Receiver: "why", Amount: 33, Memo: "test"},
	}
	blk.Height = 1
	blk.Time = 100
	
	node.BroadcastBlock(&blk)


	select {}
}
