package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ipfs/go-datastore"
	ds "github.com/ipfs/go-datastore"
	levelds "github.com/ipfs/go-ds-leveldb"
	logging "github.com/ipfs/go-log/v2"
	ldbopts "github.com/syndtr/goleveldb/leveldb/opt"
)

var log = logging.Logger("rpcserver-leveldb")

type DatastoreService struct {
	db datastore.Batching
}

func NewDatastoreService() *DatastoreService {
	//TODO: add config and options
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

func (s *DatastoreService) Get(key ds.Key) ([]byte, error) {
	log.Debugw("handle.get", "key", key)

	defer func(now time.Time) {
		log.Debugw("handled.get", "took", fmt.Sprintf("%s", time.Since(now)))
	}(time.Now())

	ctx := context.TODO()

	return s.db.Get(ctx, key)
}

func (s *DatastoreService) Has(key ds.Key) (bool, error) {
	log.Debugw("handle.has", "key", key)

	defer func(now time.Time) {
		log.Debugw("handled.has", "took", fmt.Sprintf("%s", time.Since(now)))
	}(time.Now())

	ctx := context.TODO()

	return s.db.Has(ctx, key)
}

func (s *DatastoreService) GetSize() error {
	panic("getsize - not implemented")
}

func (s *DatastoreService) Query() error {
	panic("query - not implemented")
}

func (s *DatastoreService) Put(key ds.Key, value []byte) error {
	log.Debugw("handle.put", "key", key)

	defer func(now time.Time) {
		log.Debugw("handled.put", "took", fmt.Sprintf("%s", time.Since(now)))
	}(time.Now())

	ctx := context.TODO()

	return s.db.Put(ctx, key, value)
}

func (s *DatastoreService) Delete(key ds.Key) error {
	log.Debugw("handle.delete", "key", key)

	defer func(now time.Time) {
		log.Debugw("handled.delete", "took", fmt.Sprintf("%s", time.Since(now)))
	}(time.Now())

	ctx := context.TODO()

	return s.db.Delete(ctx, key)
}

func (s *DatastoreService) Sync(prefix ds.Key) error {
	log.Debugw("handle.sync", "prefix", prefix)

	defer func(now time.Time) {
		log.Debugw("handled.sync", "took", fmt.Sprintf("%s", time.Since(now)))
	}(time.Now())

	ctx := context.TODO()

	return s.db.Sync(ctx, prefix)
}

func (s *DatastoreService) Close() error {
	log.Debugw("handle.close")

	defer func(now time.Time) {
		log.Debugw("handled.close", "took", fmt.Sprintf("%s", time.Since(now)))
	}(time.Now())

	return s.db.Close()
}

//// Batching iface

//TODO: not sure if ds.Batch will be serialzed correctly
func (s *DatastoreService) Batch() (ds.Batch, error) {
	log.Debugw("handle.batch")

	defer func(now time.Time) {
		log.Debugw("handled.batch", "took", fmt.Sprintf("%s", time.Since(now)))
	}(time.Now())

	ctx := context.TODO()

	return s.db.Batch(ctx)
}
