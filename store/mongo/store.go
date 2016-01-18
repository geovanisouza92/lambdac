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
	functions store.FunctionRepo
	runtimes  store.RuntimeRepo
}

// Indexer start the index for each repo
type Indexer interface {
	Index() error
}

func (s *mongoStore) Init(connStr string) (func(), error) {
	// logger.Println("Initializing mongo store driver.")

	// logger.Printf("Trying to connect to %q.\n", connStr)
	sess, err := mgo.DialWithTimeout(connStr, 5*time.Second)
	if err != nil {
		return nil, err
	}
	// logger.Println("Session opened.")

	sess.SetMode(mgo.Strong, true)
	sess.SetSafe(&mgo.Safe{W: 1, FSync: true})

	db := sess.DB("") // DB name came from connStr

	fn := func() {
		sess.Close()
	}

	// logger.Println("Database configured.")

	s.functions = &functionRepo{c: db.C("function")}
	if err = s.functions.(Indexer).Index(); err != nil {
		fn()
		return nil, err
	}

	// logger.Println("Functions collection configured.")

	s.runtimes = &runtimeRepo{c: db.C("runtime")}
	if err = s.runtimes.(Indexer).Index(); err != nil {
		fn()
		return nil, err
	}

	// logger.Println("Runtimes collection configured.")

	return fn, nil
}

func (s *mongoStore) Functions() store.FunctionRepo {
	return s.functions
}

func (s *mongoStore) Runtimes() store.RuntimeRepo {
	return s.runtimes
}
