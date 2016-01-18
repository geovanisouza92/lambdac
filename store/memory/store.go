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
	functions store.FunctionRepo
	runtimes  store.RuntimeRepo
}

type functionRepo struct {
	data types.Functions
}

type runtimeRepo struct {
	data types.Runtimes
}

func (s *memoryStore) Init(connStr string) (func(), error) {
	fn := func() {
	}

	s.functions = &functionRepo{data: types.Functions{}}
	s.runtimes = &runtimeRepo{data: types.Runtimes{}}

	return fn, nil
}

func (s *memoryStore) Functions() store.FunctionRepo {
	return s.functions
}

func (s *memoryStore) Runtimes() store.RuntimeRepo {
	return s.runtimes
}

func (r *functionRepo) All() (types.Functions, error) {
	return r.data, nil
}

func (r *functionRepo) FindByIDOrName(id string) (types.Function, error) {
	for _, f := range r.data {
		if f.ID == id || f.Name == id {
			return f, nil
		}
	}
	return types.Function{}, store.ErrNotFound
}

func (r *functionRepo) FindByRuntimeID(id string) (types.Functions, error) {
	out := types.Functions{}
	for _, f := range r.data {
		if f.Runtime == id {
			out = append(out, f)
		}
	}
	return out, nil
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
