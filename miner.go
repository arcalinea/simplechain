package main 

import (
    "math/rand"
    "time"
    "fmt"
)

func (node *Node) StartMiner(){
    c := make(chan string)
    go Mine(c)
    
    for {
        select{
        case sol:= <- c:
            fmt.Println(sol)
            blk := node.CreateNewBlock()
        	node.BroadcastBlock(blk)
        }
    }
}

func Mine(c chan string){
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
        rand := r.Intn(10)
        fmt.Println("Interval:", rand)
        time.Sleep(time.Duration(rand) * time.Second)
        c <- "Found a block solution"
	}
}
