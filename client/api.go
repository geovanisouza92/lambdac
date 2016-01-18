package client

import (
	"github.com/geovanisouza92/lambdac/types"
)

type API interface {
	RuntimeList() (types.Runtimes, error)
	RuntimeCreate(runtime types.Runtime) (types.Runtime, error)
	RuntimeInfo(id string) (types.Runtime, error)
	RuntimeDestroy(id string, force bool) error
}
