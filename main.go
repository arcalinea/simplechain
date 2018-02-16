package main

import (
  "fmt"
  "context"
  "time"

  "github.com/libp2p/go-libp2p"
  ipfsaddr "github.com/ipfs/go-ipfs-addr"
  peerstore "github.com/libp2p/go-libp2p-peerstore"
  floodsub "github.com/libp2p/go-floodsub"
  "os"
)

type Node struct {
    NodeID  uint64
    Mempool
    Blockchain
}


func main() {
  ctx := context.Background()

  node, err := libp2p.New(ctx, libp2p.Defaults)
  if err != nil {
      panic(err)
  }

  pubsub, err := floodsub.NewFloodSub(ctx, node)
  if err != nil {
      panic(err)
  }

  for i, addr := range node.Addrs() {
    fmt.Printf("%d: %s/ipfs/%s\n", i, addr, node.ID().Pretty())
  }

  if len(os.Args) > 1 {
      addrstr := os.Args[1]
      addr, err := ipfsaddr.ParseString(addrstr)
      if err != nil {
          panic(err)
      }
      fmt.Println("Parse Address:", addr)

      pinfo, _ := peerstore.InfoFromP2pAddr(addr.Multiaddr())

      if err := node.Connect(ctx, *pinfo); err != nil {
          fmt.Println("bootstrapping a peer failed", err)
      }
  }

  sub, err := pubsub.Subscribe("blocks")
  if err != nil {
      panic(err)
  }

  go func(){
      for range time.Tick(time.Second * 5){
          var blk Block
          blk.PrevHash = "genesis"
          blk.Transactions = []Transaction{
              {Sender: "Jay", Receiver: "Jeromy", Amount: 57, Memo: "Happy Valentine's Day <3"},
          }
          blk.Height = 1
          blk.Time = 100
          data, err := blk.Serialize()
          if err != nil {
              panic(err)
          }
          pubsub.Publish("blocks", data)
      }
  }()

  for {
      msg, err := sub.Next(ctx)
      if err != nil {
          panic(err)
      }

      blk, err := Deserialize(msg.GetData())
      if err != nil {
          panic(err)
      }

      fmt.Println(blk)

  }


  select{}
}
