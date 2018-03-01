package main

import (
	"context"
)

func main() {
	ctx := context.Background()
	
	node := CreateNewNode(ctx)
	
	var tx Transaction 
	tx.Sender = "arc"
	tx.Receiver = "why"
	tx.Amount = 1
	tx.Memo = "Hello world"
	
	node.SendTransaction(&tx)
		
	blk := node.CreateNewBlock()
	node.BroadcastBlock(blk)


	select {}
}
