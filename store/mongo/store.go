package mongo

import (
	"log"
	"os"
	"time"

	"github.com/geovanisouza92/lambdac/store"
	"gopkg.in/mgo.v2"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[mongo] ", 0)
	store.Register("mongo", new(mongoStore))
}

type mongoStore struct {
	runtimes store.RuntimeRepo
}

// Indexer start the index for each repo
type Indexer interface {
	Index() error
}

func (s *mongoStore) Init(connStr string) (func(), error) {
	sess, err := mgo.DialWithTimeout(connStr, 5*time.Second)
	if err != nil {
		return nil, err
	}

	sess.SetMode(mgo.Strong, true)
	sess.SetSafe(&mgo.Safe{W: 1, FSync: true})

	db := sess.DB("") // DB name came from connStr

	fn := func() {
		sess.Close()
	}

	s.runtimes = &runtimeRepo{c: db.C("runtime")}
	if err = s.runtimes.(Indexer).Index(); err != nil {
		fn()
		return nil, err
	}

	return fn, nil
}

func (s *mongoStore) Runtimes() store.RuntimeRepo {
	return s.runtimes
}
