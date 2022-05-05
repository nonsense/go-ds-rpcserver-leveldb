package main

import (
	"context"
	"time"

	ds "github.com/ipfs/go-datastore"
	dsq "github.com/ipfs/go-datastore/query"
)

type Datastore struct {
}

func init() {
}

var _ ds.Datastore = (*Datastore)(nil)

//var _ ds.PersistentDatastore = (*Datastore)(nil)
//var _ ds.Batching = (*Datastore)(nil)
//var _ ds.TTLDatastore = (*Datastore)(nil)
//var _ ds.GCDatastore = (*Datastore)(nil)

// NewDatastore creates a new datastore.
func NewDatastore() *Datastore {
	return &Datastore{}
}

func (d *Datastore) Put(ctx context.Context, key ds.Key, value []byte) error {
}

func (d *Datastore) Sync(ctx context.Context, prefix ds.Key) error {
}

func (d *Datastore) PutWithTTL(ctx context.Context, key ds.Key, value []byte, ttl time.Duration) error {
}

func (d *Datastore) SetTTL(ctx context.Context, key ds.Key, ttl time.Duration) error {
}

func (d *Datastore) GetExpiration(ctx context.Context, key ds.Key) (time.Time, error) {
}

func (d *Datastore) Get(ctx context.Context, key ds.Key) (value []byte, err error) {
}

func (d *Datastore) Has(ctx context.Context, key ds.Key) (bool, error) {
}

func (d *Datastore) GetSize(ctx context.Context, key ds.Key) (size int, err error) {
}

func (d *Datastore) Delete(ctx context.Context, key ds.Key) error {
}

func (d *Datastore) Query(ctx context.Context, q dsq.Query) (dsq.Results, error) {
}

func (d *Datastore) DiskUsage(ctx context.Context) (uint64, error) {
}

func (d *Datastore) Close() error {
}

func (d *Datastore) Batch(ctx context.Context) (ds.Batch, error) {
}

func (d *Datastore) CollectGarbage(ctx context.Context) (err error) {
}
