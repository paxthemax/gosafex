package store

import (
	"github.com/dgraph-io/badger"
)

// Store is the permanent storage wrapper.
type Store struct {
	dbImpl *badger.DB
}

// GlobalStore is the global permanent storage pointer.
var GlobalStore *Store

func init() {
	GlobalStore = &Store{}
}

// Open opens a store with a given cfguration.
// Returns an error if the configuration was invalid or an error was encountered.
func (s *Store) Open(cfg *Config) error {
	opts := badger.DefaultOptions
	opts.Dir = cfg.Path
	opts.ValueDir = cfg.Path

	db, err := badger.Open(opts)
	if err != nil {
		return err
	}

	s.dbImpl = db
	return nil
}

// Close closes a store.
func (s *Store) Close() {
	s.Close()
}

// Sync will perform a sync operation to the underlying storage.
// Currently sync is a no-op.
func (s *Store) Sync() {}

// NewTX creates a new transaction over the underlying storage.
func (s *Store) NewTX() TX {
	return TxData{
		impl: s.dbImpl.NewTransaction(true),
	}
}

// NewView creates a new data view over the underlying storage.
func (s *Store) NewView() View {
	return ViewData{
		impl: s.dbImpl.NewTransaction(false),
	}
}
