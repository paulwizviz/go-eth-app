package eth

import (
	"fmt"
	"net/http"
)

// RestServer is an abstraction of a RESTFul server
type RestServer struct {
	Parser Parser
}

func (r RestServer) GetCurrentBlock(w http.ResponseWriter, req *http.Request) {
	blockNum := r.Parser.GetCurrentBlock()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d", blockNum)))
}

func (r RestServer) Subscribe(w http.ResponseWriter, req *http.Request) {
	addr := req.PathValue("address")
	result := r.Parser.Subscribe(addr)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", result)))
}

func (r RestServer) GetTransactions(w http.ResponseWriter, req *http.Request) {
	addr := req.PathValue("address")
	result := r.Parser.GetTransactions(addr)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", result)))
}
