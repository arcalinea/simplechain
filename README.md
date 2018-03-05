# Simple blockchain 

A simple blockchain written in Go, using libp2p and IPFS for the networking and block storage layers. 

Blocks currently save only to memory, and chain starts from scratch each time it's run. 

## P2P networking layer

Starts a libp2p node. Sets up pubsub, subscribes to "blocks" and "transactions" topics.

Connecting to another node is currently manual. Start with first address as argument.

`./simplechain /ip4/127.0.0.1/tcp/42999/ipfs/QmdRa9h1mAxthj4ACrHULZC5yQmuiHzXDV56rWvnQaMA9o`

+ Todo: Bootstrapping and peer connection management.

## Block storage

Blocks are content-addressed and saved in memory through the IPFS data format.

## Mining 

Mining is a proof-of-work algorithm that hashes a random nonce using sha256, seeking a target solution. 

+ TODO: difficulty adjustment, save block reward.

`miner.go` also contains a function for a timeout-based mining algorithm.

## Block and transaction validation

Each block received over network is processed, and saved if it is valid.

+ TODO: Reorg behavior

## Wallet 

Wallet and account behavior is a stub, currently only contains a function that returns an arbitrary random string as a "pubkey".

+ TODO: Key mgmt for accounts, set and track balances.

## RPC interface

Starts up an http server, provides command line RPC interface. Cli must be built and run separately.

`go build -o simple-cli ./cli`

`./simple-cli getinfo`

### Commands
- `getinfo` - Takes no arguments, returns blockchain info.
- `sendtx` - Takes receiver address, amount, and optional values of sender address and a memo. 
    - `./cli sendtx you 100 -from=me -memo=hi`
