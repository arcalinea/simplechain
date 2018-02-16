
P2P networking layer:
----------------------
Start a libp2p node.
Potentially connect to another libp2p node.

Set up pubsub.
Subscribe to "blocks" topic.
(Subscribe to "txs" topic, "alerts", etc)

peer connection mgmt

Block validation:
------------------
Process each block received, check if good.
If good, accept it.
Log block to debug log.
Save block to disk.

RPC interface:
---------------
Start up an http server, for json-rpc interface.
Build out regtest mode.

`getinfo`

mining:
`generate` command for new blocks.

wallet:
`getnewaddress`
`getbalance`
`sendtransaction`
`listtransactions`
