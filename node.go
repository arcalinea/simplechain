package main

type Node struct {
	NodeID uint64
	mempool *Mempool
	blockchain *Blockchain
}

func (node *Node) SendTransaction(tx *Transaction) string {
    
}
