package main

import (
	"flag"
	"net/http"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gorilla/mux"
	logging "github.com/ipfs/go-log/v2"
)

var repopath string

func init() {
	logging.SetLogLevel("*", "debug")

	flag.StringVar(&repopath, "repopath", "", "path for repo")
}

func main() {
	flag.Parse()

	ds := NewDatastoreService(repopath)
	server := rpc.NewServer()
	server.RegisterName("rpcdatastore", ds)

	router := mux.NewRouter()
	router.Handle("/", server)

	log.Debugw("leveldb.path", "path", repopath)
	log.Infow("http.listen", "server is listening", "localhost:8089")
	http.ListenAndServe("localhost:8089", router)
}
