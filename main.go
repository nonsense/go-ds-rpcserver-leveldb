package main

import (
	"net/http"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gorilla/mux"
	logging "github.com/ipfs/go-log/v2"
)

func init() {
	logging.SetLogLevel("*", "debug")
}

func main() {
	ds := NewDatastoreService()
	server := rpc.NewServer()
	server.RegisterName("rpcdatastore", ds)

	router := mux.NewRouter()
	router.Handle("/", server)
	http.ListenAndServe("localhost:8089", router)
}
