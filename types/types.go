package types

type GetInfoResponse struct {
    BlockHeight     uint64
}   

type SendTxResponse struct {
    Txid       string
}

type GetNewAddressResponse struct {
    Address     string
}
