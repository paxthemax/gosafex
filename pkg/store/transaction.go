package store

import "github.com/dgraph-io/badger"

// TxData contrains all data for a transaction on the store.
type TxData struct {
	impl *badger.Txn
}

// Commit will commit any changes to the underlying storage.
func (tx TxData) Commit() (err error) {
	err = tx.impl.Commit(nil)
	tx.impl.Discard()
	return
}

// Rollback will roll back any changes in the transaction.
func (tx TxData) Rollback() {
	panic("not implemented")
}

// Sync will perform a sync operation with the underlying storage.
// Currently sync is a no-op.
func (tx TxData) Sync() {}

// LoadValue will look up and return a value by a give key form the transaction.
// The value returned is the last changed value in the transaction.
// Returns an error if the key was not found.
func (tx TxData) LoadValue(k Key) (v Value, err error) {
	entry, err := tx.impl.Get(k)
	if err != nil {
		return nil, err
	}
	data, err := entry.ValueCopy(nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// StoreValue will store the new value of the given key.
// Value is stored in the trasaction. Permanent changes to the store must be synced via Commit.
func (tx TxData) StoreValue(k Key, v Value) (err error) {
	// TODO: store to the cache first.
	buf := make([]byte, len(v))
	copy(buf, v)
	return tx.impl.Set(k, buf)
}
