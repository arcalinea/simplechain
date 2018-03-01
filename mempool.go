package main 

type Mempool struct {
    transactions []Transaction
}

func (mempool *Mempool) SelectTransactions() []Transaction {
    return mempool.transactions
}
