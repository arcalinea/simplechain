package main

import (
	"context"
)

func main() {
	ctx := context.Background()
	
	node := CreateNewNode(ctx)
	
	var blk Block
	blk.Transactions = []Transaction{
		{Sender: "arc", Receiver: "why", Amount: 57, Memo: "test"},
	}
	blk.Height = 1
	blk.Time = 100
	
	node.BroadcastBlock(&blk)


	select {}
}
