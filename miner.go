package main 

import (
    "math/rand"
    "time"
    "fmt"
)

func (node *Node) StartMiner(){
    c := make(chan string)
    go node.Mine(c)
}

func (node *Node) Mine(c chan string){
    go FindSols(c)
    for {
        select{
        case sol:= <- c:
            fmt.Println(sol)
            blk := node.CreateNewBlock()
            node.BroadcastBlock(blk)
        }
    }
}

func FindSols(c chan string){
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
        rand := r.Intn(10)
        fmt.Println("Interval:", rand)
        time.Sleep(time.Duration(rand) * time.Second)
        c <- "Found a block solution"
	}
}
