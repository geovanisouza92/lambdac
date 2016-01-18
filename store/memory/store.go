// Dummy implementation of store -- DO NOT USE THIS IN PRODUCTION
package memory

import (
	"github.com/geovanisouza92/lambdac/store"
	"github.com/geovanisouza92/lambdac/types"
)

func init() {
	store.Register("memory", new(memoryStore))
}

type memoryStore struct {
	runtimes store.RuntimeRepo
}

type runtimeRepo struct {
	data types.Runtimes
}

func (s *memoryStore) Init(connStr string) (func(), error) {
	fn := func() {
	}

	s.runtimes = &runtimeRepo{data: types.Runtimes{}}

	return fn, nil
}

func (s *memoryStore) Runtimes() store.RuntimeRepo {
	return s.runtimes
}

func (r *runtimeRepo) All() (types.Runtimes, error) {
	return r.data, nil
}

func (r *runtimeRepo) FindByIDOrName(id string) (types.Runtime, error) {
	for _, r := range r.data {
		if r.ID == id || r.Name == id {
			return r, nil
		}
	}
	return types.Runtime{}, nil
}

func (r *runtimeRepo) Remove(id string) error {
	for i, rt := range r.data {
		if rt.ID == id {
			r.data = append(r.data[i:], r.data[i+1:]...)
			return nil
		}
	}
	return store.ErrNotFound
}

func (r *runtimeRepo) Create(runtime types.Runtime) (types.Runtime, error) {
	r.data = append(r.data, runtime)
	return runtime, nil
}
