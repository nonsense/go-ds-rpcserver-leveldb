package main

import (
	ds "github.com/ipfs/go-datastore"
	dsq "github.com/ipfs/go-datastore/query"
)

type GetRequest struct {
	Key ds.Key
}

type GetResponse struct {
	Value []byte
	Error error
}

type PutRequest struct {
	Key   ds.Key
	Value []byte
}

type PutResponse struct {
	Error error
}

type HasRequest struct {
	Key ds.Key
}

type HasResponse struct {
	Exists bool
	Error  error
}

type QueryRequest struct {
	Query dsq.Query
}

type QueryResponse struct {
	Results dsq.Results
	Error   error
}

type DeleteRequest struct {
}

type DeleteResponse struct {
}

type GetSizeRequest struct {
}

type GetSizeResponse struct {
}

type SyncRequest struct {
}

type SyncResponse struct {
}

type CloseRequest struct {
}

type CloseResponse struct {
}

type BatchRequest struct {
}

type BatchResponse struct {
}
