package main

import (
	"context"
	"fmt"
	"strconv"
	"net/http"
)

func main() {
	ctx := context.Background()
	
	node := CreateNewNode(ctx)
	
	node.StartMiner()
		
	http.HandleFunc("/sendtx", func(w http.ResponseWriter, r *http.Request) {
		from := r.FormValue("from")
		to := r.FormValue("to")
		amount := r.FormValue("amount")
		memo := r.FormValue("memo")
		
		fmt.Println("call sendtx", from, to, amount, memo)
		
		amt, err := strconv.ParseInt(amount, 10, 64)
		if err != nil {
			panic(err)
		}
	
		tx := Transaction{
			Sender: from,
			Receiver: to,
			Amount: uint64(amt),
			Memo: memo, 
		}
		
		node.SendTransaction(&tx)
	})
	

	http.ListenAndServe(":1234", nil)
	
}
