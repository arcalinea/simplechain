package main

import (
	"context"
	// "fmt"
	// "net/http"
)

func main() {
	ctx := context.Background()
	
	node := CreateNewNode(ctx)
	
	node.StartMiner()
	
	var tx Transaction 
	tx.Sender = "arc"
	tx.Receiver = "why"
	tx.Amount = 1
	tx.Memo = "Hello world"
	
	node.SendTransaction(&tx)
		
	
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(r.URL)
	// 	w.Write([]byte("hello"))
	// })
	// 
	// 
	// http.ListenAndServe(":1234", nil)
	
}
