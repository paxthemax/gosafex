package store

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
	// We can only accress valid iterators.
	Valid() bool

	// Next advances the iterator.
	// Iterator moves to the next sequential key in the order of iteration.
	// Returns error if out of scope.
	Next() error

	// Key returns the key of the cursor.
	// Returns error if cursor position is invalid.
	Key() (key []byte, err error)

	// Value returns the value of the cursor.
	// Returns error if cursor position is invalid.
	Value() (value []byte, err error)

	// Close releases the iterator.
	Close()
}

// ReverseIterator is an Iterator that can also move the cursor in reverse.
type ReverseIterator interface {
	Iterator
	// Prev moves the cursor in reverse.
	// Iterator moves to the previous sequential key in the order of iteration.
	// Returns error if out of scope.
	Prev() error
}

// KVIterator is the iterator over a slice key-value pairs.
type KVIterator struct {
	data []KV
	idx  int
}

// NewKVIterator creates a new iterator over a given key-value pair slice.
func NewKVIterator(data []KV) *KVIterator {
	return &KVIterator{data: data}
}

func (it *KVIterator) idxInBounds(pos int) bool { return pos < len(it.data) }

func (it *KVIterator) mustBeInBounds(pos int) error {
	if ok := it.idxInBounds(it.idx + 1); !ok {
		return ErrOutOfBounds
	}
	return nil
}

func (it *KVIterator) value() []byte {
	return it.data[it.idx].Value
}

func (it *KVIterator) key() []byte {
	return it.data[it.idx].Value
}

// Valid implements Iterator and returns true if iterator is valid.
func (it *KVIterator) Valid() bool { return it.idxInBounds(it.idx) }

// Next implements Iterator and advances the cursor.
func (it *KVIterator) Next() error {
	if err := it.mustBeInBounds(it.idx + 1); err != nil {
		return ErrOutOfBounds
	}
	it.idx++
	return nil
}

// Key implements Iterator and returns the key of the cursor.
func (it *KVIterator) Key() (key []byte, err error) {
	if err = it.mustBeInBounds(it.idx); err != nil {
		return nil, err
	}
	return it.value(), nil
}

// Value implements Iterator and returns the value of the cursor.
func (it *KVIterator) Value() (val []byte, err error) {
	if err = it.mustBeInBounds(it.idx); err != nil {
		return nil, err
	}
	return it.value(), nil
}

// Close implements Iterator and releases the current iterator.
func (it *KVIterator) Close() {
	it.data = nil
}

// Prev implements Iterator and moves the cursor back.
func (it *KVIterator) Prev() error {
	if err := it.mustBeInBounds(it.idx - 1); err != nil {
		return ErrOutOfBounds
	}
	it.idx--
	return nil
}
