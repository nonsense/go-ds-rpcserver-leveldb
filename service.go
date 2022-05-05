package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ipfs/go-datastore"
	levelds "github.com/ipfs/go-ds-leveldb"
	logging "github.com/ipfs/go-log/v2"
	ldbopts "github.com/syndtr/goleveldb/leveldb/opt"
)

var log = logging.Logger("rpcserver-badger")

type DatastoreService struct {
	db datastore.Batching
}

func NewDatastoreService() *DatastoreService {
	db, err := levelDs("/tmp/smth", false)
	if err != nil {
		panic(err)
	}

	return &DatastoreService{
		db: db,
	}
}

func levelDs(path string, readonly bool) (datastore.Batching, error) {
	return levelds.NewDatastore(path, &levelds.Options{
		Compression: ldbopts.NoCompression,
		NoSync:      false,
		Strict:      ldbopts.StrictAll,
		ReadOnly:    readonly,
	})
}

func (s *DatastoreService) Get(r *http.Request, req *GetRequest, res *GetResponse) error {
	log.Debugw("handle.get", "key", req.Key)

	defer func(now time.Time) {
		log.Debugw("handled.get", "took", fmt.Sprintf("%s", time.Since(now)))
	}(time.Now())

	ctx := r.Context()

	value, err := s.db.Get(ctx, req.Key)
	res = &GetResponse{
		Value: value,
		Error: err,
	}

	return nil
}

func (s *DatastoreService) Has(r *http.Request, req *HasRequest, res *HasResponse) error {
	log.Debugw("handle.has", "key", req.Key)

	defer func(now time.Time) {
		log.Debugw("handled.has", "took", fmt.Sprintf("%s", time.Since(now)))
	}(time.Now())

	ctx := r.Context()

	exists, err := s.db.Has(ctx, req.Key)
	res = &HasResponse{
		Exists: exists,
		Error:  err,
	}

	return nil
}

func (s *DatastoreService) GetSize(r *http.Request, req *GetSizeRequest, res *GetSizeResponse) error {
	panic("getsize - not implemented")
}

func (s *DatastoreService) Query(r *http.Request, req *QueryRequest, res *QueryResponse) error {
	log.Debugw("handle.query")

	defer func(now time.Time) {
		log.Debugw("handled.query", "took", fmt.Sprintf("%s", time.Since(now)))
	}(time.Now())

	ctx := r.Context()

	results, err := s.db.Query(ctx, req.Query)
	res = &QueryResponse{
		Results: results,
		Error:   err,
	}

	return nil
}

func (s *DatastoreService) Put(r *http.Request, req *PutRequest, res *PutResponse) error {
	log.Debugw("handle.put", "key", req.Key)

	defer func(now time.Time) {
		log.Debugw("handled.put", "took", fmt.Sprintf("%s", time.Since(now)))
	}(time.Now())

	ctx := r.Context()

	err := s.db.Put(ctx, req.Key, req.Value)
	res = &PutResponse{
		Error: err,
	}

	return nil
}

func (s *DatastoreService) Delete(r *http.Request, req *DeleteRequest, res *DeleteResponse) error {
	panic("delete - not implemented")
}

func (s *DatastoreService) Sync(r *http.Request, req *SyncRequest, res *SyncResponse) error {
	panic("sync - not implemented")
}

func (s *DatastoreService) Close(r *http.Request, req *CloseRequest, res *CloseResponse) error {
	panic("close - not implemented")
}

func (s *DatastoreService) Batch(r *http.Request, req *BatchRequest, res *BatchResponse) error {
	panic("batch - not implemented")
}
