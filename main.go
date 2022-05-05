package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"

	logging "github.com/ipfs/go-log/v2"
)

func init() {
	logging.SetLogLevel("*", "debug")
}

func main() {
	rpcServer := rpc.NewServer()

	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	ds := NewDatastoreService()

	err := rpcServer.RegisterService(ds, "")
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.Handle("/", rpcServer)

	fmt.Println("Listening...")
	http.ListenAndServe(":8089", router)
}
