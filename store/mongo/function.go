package mongo

import (
	"github.com/geovanisouza92/lambdac/store"
	"github.com/geovanisouza92/lambdac/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type functionRepo struct {
	c *mgo.Collection
}

func (r *functionRepo) Index() (err error) {
	idIdx := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
		Name:       "idx_function_id",
	}
	err = r.c.EnsureIndex(idIdx)
	if err != nil {
		return
	}
	nameIdx := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
		Name:       "idx_function_name",
	}
	err = r.c.EnsureIndex(nameIdx)
	if err != nil {
		return
	}
	runtimeIdx := mgo.Index{
		Key:        []string{"runtime"},
		Background: true,
		Sparse:     true,
		Name:       "idx_function_runtime",
	}
	err = r.c.EnsureIndex(runtimeIdx)
	return
}

func (r *functionRepo) All() (functions types.Functions, err error) {
	q := r.c.Find(nil)
	err = q.All(&functions)
	if err == mgo.ErrNotFound {
		err = store.ErrNotFound
	}
	// TODO Handle other errors
	return
}

func (r *functionRepo) FindByIDOrName(id string) (function types.Function, err error) {
	q := r.c.Find(bson.M{
		"$or": []interface{}{
			bson.M{"id": id},
			bson.M{"name": id},
		},
	})
	err = q.One(&function)
	if err == mgo.ErrNotFound {
		err = store.ErrNotFound
	}
	// TODO Handle other errors
	return
}

func (r *functionRepo) FindByRuntimeID(id string) (functions types.Functions, err error) {
	q := r.c.Find(bson.M{"runtime": id})
	err = q.All(&functions)
	if err == mgo.ErrNotFound {
		// Ignore ErrNotFound
		err = nil
		functions = types.Functions{}
	}
	// TODO Handle other errors
	return
}

func (r *functionRepo) Create(function types.Function) (out types.Function, err error) {
	err = r.c.Insert(function)
	out = function
	return
}

func (r *functionRepo) Update(function types.Function) (err error) {
	q := bson.M{"id": function.ID}
	err = r.c.Update(q, function)
	return
}

func (r *functionRepo) Remove(id string) (err error) {
	err = r.c.Remove(bson.M{"id": id})
	if err != nil {
		err = store.ErrNotFound
	}
	// TODO Handle other errors
	return
}
