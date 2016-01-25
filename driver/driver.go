// Runtime driver abstraction
package driver

import (
	"fmt"
	"sync"

	"github.com/geovanisouza92/lambdac/types"
)

var (
	driversMu sync.Mutex
	drivers   = make(map[string]Driver)
)

type Driver interface {
	// Configure driver
	Init(options []string) error
	// Create a function instance
	Create(function types.Function, runtime types.Runtime) (string, error)
	// Start a function instance
	Start(id string) error
	// Stop a function instance
	Stop(id string) error
	// Destroy a function instance
	Destroy(id string) error
}

func Register(name string, driver Driver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("driver: Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("driver: Register called twice for driver " + name)
	}
	drivers[name] = driver
}

func Open(name string) (Driver, error) {
	driversMu.Lock()
	driver, ok := drivers[name]
	driversMu.Unlock()
	if !ok {
		return nil, fmt.Errorf("driver: unknown driver %q (forgotten import?)", name)
	}
	return driver, nil
}
