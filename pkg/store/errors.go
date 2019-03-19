package store

import "github.com/pkg/errors"

// ErrKeyNotFound returns an error that a key was not found.
func ErrKeyNotFound(err error) error { return errors.Wrap(err, "Key not found") }

// ErrOperationFailed returns an error that a generic store operation fails.
func ErrOperationFailed(err error) error { return errors.Wrap(err, "Operation failed") }

// ErrIteration returns an error that moving the iterator failed.
func ErrIteration(err error) error { return errors.Wrap(err, "Moving iterator failed") }
