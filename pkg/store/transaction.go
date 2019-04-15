package store

// BaseTransaction can commit write changes, revert write changes or sync with the underlying storage.
type BaseTransaction interface {
	Commit() error
	Rollback()
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
	BaseTransaction
	ValueLoader
	ValueStorer
}

// TX is an alias for ValueTransaction.
type TX = ValueTransaction
