package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/paulwizviz/go-eth-app/internal/eth"
)

// RestServer is an abstraction of a RESTFul server
type RestServer struct {
	Parser eth.Parser
}

type GetCurrentBlockResponse struct {
	Block     string `json:"block"`
	Addresses string `json:"addresses"`
}

func (r RestServer) GetCurrentBlock(w http.ResponseWriter, req *http.Request) {
	blockNum := r.Parser.GetCurrentBlock()

	resp := GetCurrentBlockResponse{
		Block:     blockNum,
		Addresses: fmt.Sprintf("http://%s/addresses", req.Host),
	}
	json, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(json))
}

func (r RestServer) Subscribe(w http.ResponseWriter, req *http.Request) {
	addr := req.PathValue("address")

	sub := r.Parser.Subscribe(addr)

	log.Printf("New subscription for %s; ID: %s\n", addr, sub.ID)

	// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.(http.Flusher).Flush()

outer:
	for {
		select {
		case txn := <-sub.Ch:
			fmt.Fprintf(w, "data: %s\n\n", string(txn))
			w.(http.Flusher).Flush()

		case <-req.Context().Done():
			sub.Unsubscribe()
			log.Printf("Unsubscribing %s", sub.ID)
			break outer
		}
	}
}

type GetTransactionsLinks struct {
	Addresses string `json:"addresses"`
	Subscribe string `json:"subscribe"`
}
type GetTransactionsResponse struct {
	Links        GetTransactionsLinks `json:"links"`
	Transactions []eth.Transaction    `json:"transactions"`
}

func (r RestServer) GetTransactions(w http.ResponseWriter, req *http.Request) {
	addr := req.PathValue("address")
	result := r.Parser.GetTransactions(addr)
	w.WriteHeader(http.StatusOK)

	resp := GetTransactionsResponse{
		Transactions: result,
		Links: GetTransactionsLinks{
			Addresses: fmt.Sprintf("http://%s/addresses", req.Host),
			Subscribe: fmt.Sprintf("http://%s/addresses/%s/subscribe", req.Host, addr),
		},
	}
	json, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(json))
}

type Address struct {
	Address         string `json:"address"`
	TransactionsURL string `json:"transactions"`
	Count           int64  `json:"count"`
}
type GetAddressesResponse struct {
	Addresses []Address `json:"addresses"`
}

func (r RestServer) GetAddresses(w http.ResponseWriter, req *http.Request) {
	keys := r.Parser.GetAddresses()
	w.WriteHeader(http.StatusOK)

	addresses := []Address{}
	for _, k := range keys {
		address := Address{
			Address:         k,
			TransactionsURL: fmt.Sprintf("http://%s/addresses/%s", req.Host, k),
			Count:           r.Parser.GetCount(k),
		}
		addresses = append(addresses, address)
	}

	sort.Slice(addresses, func(i, j int) bool {
		return addresses[i].Count > addresses[j].Count
	})

	resp := GetAddressesResponse{addresses}
	json, err := json.Marshal(&resp)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(json))
}
