package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/ipfs/go-datastore"
	ds "github.com/ipfs/go-datastore"
	dsq "github.com/ipfs/go-datastore/query"
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

	dir, err := ioutil.TempDir("", "ds-leveldb")
	if err != nil {
		panic(err)
	}

	db, err := levelDs(dir, false)
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

	value, err := s.db.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	log.Debugw("handle.got", "key", key, "value", string(value))

	return value, nil
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

//func (s *DatastoreService) Query(q dsq.Query) (dsq.Results, error) {
func (s *DatastoreService) Query(q dsq.Query) ([]dsq.Entry, error) {
	log.Debugw("handle.query", "query", q)

	defer func(now time.Time) {
		log.Debugw("handled.query", "took", fmt.Sprintf("%s", time.Since(now)))
	}(time.Now())

	ctx := context.TODO()

	results, err := s.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	entries, err := results.Rest()
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (s *DatastoreService) Put(key ds.Key, value []byte) error {
	log.Debugw("handle.put", "key", key, "value", string(value))

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
