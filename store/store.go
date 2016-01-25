// Data storage abstraction
package store

import (
	"errors"
	"fmt"
	"sync"

	"github.com/geovanisouza92/lambdac/types"
)

var (
	ErrNotFound = errors.New("Not found")

	storesMu sync.Mutex
	stores   = make(map[string]Store)
)

// Store is the interface used to interact with a data store
type Store interface {
	// Init the store (open files, connections, allocate memory, etc.)
	Init(connStr string) (func(), error)

	// Functions gives access to the underlying FunctionRepo instance
	Functions() FunctionRepo

	// Runtimes gives access to the underlying RuntimeRepo instance
	Runtimes() RuntimeRepo
}

// FunctionRepo is the interface that holds function information
type FunctionRepo interface {
	// All get functions matching criteria
	All() (types.Functions, error)

	// FindByIDOrName finds a function by ID or name (can be partial)
	FindByIDOrName(id string) (types.Function, error)

	// FindByRuntimeID finds all functions matching with specified runtime ID
	FindByRuntimeID(id string) (types.Functions, error)

	// Create a new function
	Create(function types.Function) (types.Function, error)

	// Update a function
	Update(function types.Function) error

	// Remove a function
	Remove(id string) error
}

// RuntimeRepo is the interface that holds runtime information
type RuntimeRepo interface {
	// All get runtimes matching criteria
	All() (types.Runtimes, error)

	// FindByIDOrName finds a runtime by ID or name (can be partial)
	FindByIDOrName(id string) (types.Runtime, error)

	// Create a new runtime
	Create(runtime types.Runtime) (types.Runtime, error)

	// Remove a runtime
	Remove(id string) error
}

// Register a new store
func Register(name string, store Store) {
	storesMu.Lock()
	defer storesMu.Unlock()
	if store == nil {
		panic("store: Register store is nil")
	}
	if _, dup := stores[name]; dup {
		panic("store: Register called twice for driver " + name)
	}
	stores[name] = store
}

// Open (find) a store instance by name
func Open(name string) (Store, error) {
	storesMu.Lock()
	store, ok := stores[name]
	storesMu.Unlock()
	if !ok {
		return nil, fmt.Errorf("store: unknown store %q (forgotten import?)", name)
	}
	return store, nil
}
