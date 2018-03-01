package main 

type Mempool struct {
    transactions []Transaction
}

func (mempool *Mempool) AddTx(tx *Transaction) {
    mempool.transactions = append(mempool.transactions, *tx)
}

func (mempool *Mempool) SelectTransactions() []Transaction {
    return mempool.transactions
}
