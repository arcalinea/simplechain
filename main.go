package main

import (
	"context"
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
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
		
		err = json.NewEncoder(w).Encode(node.SendTransaction(&tx))
		if err != nil {
			panic(err)
		}
	})
	
	http.HandleFunc("/getinfo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Call getinfo")
		
		err := json.NewEncoder(w).Encode(node.GetInfo())
		// res, err := json.Marshal(node.GetInfo())
		if err != nil {
			panic(err)
		}
		// fmt.Fprintf(w, string(res))
	})
	

	http.ListenAndServe(":1234", nil)
	
}
