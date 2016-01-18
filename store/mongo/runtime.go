package mongo

import (
	"github.com/geovanisouza92/lambdac/store"
	"github.com/geovanisouza92/lambdac/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type runtimeRepo struct {
	c *mgo.Collection
}

func (r *runtimeRepo) Index() error {
	nameIdx := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
		Name:       "idx_runtime_name",
	}
	return r.c.EnsureIndex(nameIdx)
}

func (r *runtimeRepo) All() (runtimes types.Runtimes, err error) {
	q := r.c.Find(nil)
	err = q.All(&runtimes)
	return
}

func (r *runtimeRepo) FindByIDOrName(id string) (runtime types.Runtime, err error) {
	q := r.c.Find(bson.M{
		"$or": []interface{}{
			bson.M{
				"id": bson.M{
					"$regex": bson.RegEx{".*" + id + ".*", ""},
				},
			},
			bson.M{
				"name": bson.M{
					"$regex": bson.RegEx{".*" + id + ".*", ""},
				},
			},
		},
	})
	err = q.One(&runtime)
	if err == mgo.ErrNotFound {
		err = store.ErrNotFound
	}
	// TODO Handle other errors
	return
}

func (r *runtimeRepo) Create(runtime types.Runtime) (out types.Runtime, err error) {
	err = r.c.Insert(runtime)
	out = runtime
	return
}

func (r *runtimeRepo) Remove(id string) (err error) {
	err = r.c.Remove(bson.M{"id": id})
	if err != nil {
		err = store.ErrNotFound
	}
	// TODO Handle other errors
	return
}
