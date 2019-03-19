package store

// ReadChecker is a simple interface to query data.
type ReadChecker interface {
	// Get returns the value of a given key.
	// Returns error if key does not exists.
	// Panics if key does not exist.
	Get([]byte) ([]byte, error)

	// Has returns true if a given key exists.
	// Panics if key is invalid.
	Has([]byte) bool
}

// WriteDeleter is a simple interface to write and delete data.
type WriteDeleter interface {
	// Write sets the value of a given key.
	// Returns error if the key was not found or if the operation failed.
	Write(key []byte, val []byte) error

	// Delete removes a single key.
	// Returns error if the key was not found or if the operation failed.
	Delete(key []byte) error
}

// Store is the interface all stores should implement.
// It requires basic CRUD operations and creating transactions.
type Store interface {
	ReadChecker
	WriteDeleter

	// NewTransaction creates a new Transaction.
	// Transactions can be used to write several operations atomically.
	NewTransaction() Transaction
}

// Transaction is a set of atomic changes that are set to be committed to the underlying store.
type Transaction interface {
	WriteDeleter

	// Commit writes all transaction changes to the underlying store.
	// Returns error on operation failed.
	Commit() error
}

/*
Iterator allows access to a set of items within a range of keys.
These items may all be preloaded, or loaded on demand.
  Usage:
  var itr Iterator = ...
  defer itr.Close()
  for ; itr.Valid(); itr.Next() {
    k, v := itr.Key(); itr.Value()
    ...
  }
*/
type Iterator interface {
	// Valid returns if the cursor position is valid.
	// Once the iterator becomes invalid, it stays invalid.
	Valid() bool

	// Next advances the iterator.
	// IIterator moves to the next sequential key in the order of iteration.
	// Returns error on failure.
	Next() error

	// Key returns the key of the cursor.
	// Returns error if cursor position is invalid.
	Key() (key []byte, err error)

	// Value returns the value of the cursor.
	// Returns error if cursor position is invalid.
	Value() (value []byte, err error)

	// Close releases the Iterator.
	Close()
}
