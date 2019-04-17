package store

import "github.com/dgraph-io/badger"

// ViewData contains all data for a data view over the store.
type ViewData struct {
	impl *badger.Txn
}

// LoadValue will look up and return a value by a give key form the view.
// Returns an error if the key was not found.
func (view ViewData) LoadValue(k Key) (v Value, err error) {
	entry, err := view.impl.Get(k)
	if err != nil {
		return nil, err
	}
	data, err := entry.ValueCopy(nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Sync will perform a sync operation with the underlying storage.
// Currently sync is a no-op.
func (view ViewData) Sync() {}
