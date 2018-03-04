# Simple blockchain 

A simple blockchain written in Go, using libp2p and IPFS for the networking and block storage layers. 

Blocks currently save only to memory, and chain starts from scratch each time it's run. 

## P2P networking layer

Starts a libp2p node. Sets up pubsub, subscribe to "blocks" and "transactions" topics.

Connecting to another node is currently manual. Start with first address as argument.

Todo: Bootstrapping and peer connection management.

## Block storage

Blocks are content-addressed and saved in memory through the IPFS data format.

## Mining 

Mining is a proof-of-work algorithm that hashes a random nonce using sha256, seeking a target solution. Currently has no difficulty adjustment.

`miner.go` also contains a function for a timeout-based mining algorithm, to be used to reduce load when testing, or when nodes are trusted.

## Block and transaction validation

Each block received over network is processed, and saved if it is valid.

# Wallet 

Wallet and account behavior is a stub, currently only contains a function that returns an arbitrary random string as a "pubkey".

## RPC interface

Starts up an http server, which can be interacted with through the command line.

Blockchain:
`getinfo` - Takes no arguments, returns blockchain info.

Wallet:
`sendtx` - Takes receiver address, amount, and optional values of sender address and a memo. `./cli sendtx you 100 -from=me -memo=hi`
