package store

import (
	"github.com/dgraph-io/badger"
)

// Storer can store data to a permanent storage.
// Storer can open the permanent storage. Storer can close the permanent storage and release resources.
// Storer can create a new transaction.
type Storer interface {
	Open(cfg *Config) error
	Close()
	NewTX() *TX
}

// Store is the permanent storage wrapper.
type Store struct {
	db *badger.DB
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

	s.db = db
	return nil
}
