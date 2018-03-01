package main

import (
	"context"
)

func main() {
	ctx := context.Background()
	
	node := CreateNewNode(ctx)
	
	blk := node.CreateNewBlock()
	
	node.BroadcastBlock(blk)


	select {}
}
