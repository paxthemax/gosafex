package store

// Empty store persists no data.
// Emptys can be wrapped by caching layers.
type Empty struct{}

// NewEmpty creates a new empty store.
func NewEmpty() *Empty { return &Empty{} }

// Get always returns nil
func (e Empty) Get(key []byte) []byte { return nil }

// Has always returns false
func (e Empty) Has(key []byte) bool { return false }

// Set is a noop
func (e Empty) Set(key, value []byte) {}

// Delete is a noop
func (e Empty) Delete(key []byte) {}

// Iterator is always empty
func (e Empty) Iterator(start, end []byte) Iterator {
	// TODO: return empty iterator
	return nil
}
