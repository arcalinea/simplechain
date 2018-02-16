package main


type Transaction struct {
    Sender string
    Receiver string
    Amount uint64
    Memo string
}

func (tx *Transaction) Serialize() ([]byte, error){
    return json.Marshal(tx)
}

func Deserialize(buf []byte) (*Transaction, error){
    var tx Transaction
    err := json.Unmarshal(buf, &tx)
    if err != nil {
        return nil, err
    }
    return &tx, nil
}
