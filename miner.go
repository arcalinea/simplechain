package main 

import (
    "math/rand"
    "time"
    "fmt"
    "crypto/sha256"
    "encoding/base64"
    // "math/big"
)

func (node *Node) StartMiner(){
    c := make(chan *Block)
    go node.Mine(c)
}

func (node *Node) Mine(c chan *Block){
    go node.FindSolsHash(c)
    for {
        select{
        case blk := <- c:
            fmt.Println(blk)
            node.BroadcastBlock(blk)
        }
    }
}

func FindSolsTimeout(c chan string){
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
        rand := r.Intn(10) // Adjusts variance of block speed
        fmt.Println("Interval:", rand)
        time.Sleep(time.Duration(rand) * time.Second)
        c <- "Found a block solution"
	}
}

func isWinner(ticket string) bool {
    // num := big.NewInt(0).SetBytes(ticket)
    winner := "000"
    res := false
    if ticket[:len(winner)] == winner {
        return true
    }
    return res
}

func (node *Node) FindSolsHash(c chan *Block){
    blk := node.CreateNewBlock()
    blk.Nonce = make([]byte, 32)
    for {
        rand.Read(blk.Nonce)
        guess := sha256.Sum256(blk.Serialize())
        ticket := base64.StdEncoding.EncodeToString(guess[:])
        if isWinner(ticket) {
            fmt.Println("Ticket:", ticket)
            c <- blk
        } 
    }
}
