package main 

type Mempool struct {
    transactions map[[32]byte]Transaction
}

func NewMempool() *Mempool{
    return &Mempool{
        transactions: make(map[[32]byte]Transaction),
    }
}

func (mempool *Mempool) AddTx(tx *Transaction) {
    txid := tx.GetTxid()
    mempool.transactions[txid] = *tx
}

func (mempool *Mempool) removeTxs(txs []Transaction){
    for _, tx := range txs {
        txid := tx.GetTxid()
        delete(mempool.transactions, txid)
    }
}

func (mempool *Mempool) SelectTransactions() []Transaction {
    var txs []Transaction
    for _, v := range mempool.transactions {
        txs = append(txs, v)
    }
    return txs
}
