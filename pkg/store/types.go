package store

// Key is the key for a key-value map.
type Key = []byte

// Value is the value in a key-value map.
type Value = []byte

// Commiter can commit its changes to an underlying storeage.
type Commiter interface {
	Commit() error
}

// Rollbacker can roll back changes it made to itself.
type Rollbacker interface {
	Rollback()
}

// Syncer can sync its state to match the underlying storage.
type Syncer interface {
	Sync()
}

// ValueLoader can load a Value.
type ValueLoader interface {
	LoadValue(k Key) (v Value, err error)
}

// ValueStorer can store a Value.
type ValueStorer interface {
	StoreValue(k Key, v Value) (err error)
}

// ValueTransaction is a transaction that can load and store values.
type ValueTransaction interface {
	Commiter
	Rollbacker
	Syncer
	ValueLoader
	ValueStorer
}

// ValueView is a read-only data view.
type ValueView interface {
	Syncer
	ValueLoader
}

// TX is an alias for ValueTransaction.
type TX = ValueTransaction

// View is an alias for ValueView.
type View = ValueView

// StoreWrapper is a wrapper around permanent storage.
type StoreWrapper interface {
	Open(cfg *Config) error
	Close()
	NewTX() TX
	NewView() View
}
