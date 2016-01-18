// LambdaC API server
package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/geovanisouza92/lambdac/store"
	_ "github.com/geovanisouza92/lambdac/store/memory"
	_ "github.com/geovanisouza92/lambdac/store/mongo"
	"github.com/gorilla/mux"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[server] ", 0)
}

type Server struct {
	r *mux.Router // Router
	s store.Store // Data store
	f func()      // Function to close store
}

func New(storeName, connString string) (*Server, error) {
	// Load store
	st, err := store.Open(storeName)
	if err != nil {
		return nil, err
	}

	// Init store
	fn, err := st.Init(connString)
	if err != nil { // Store must execute fn()
		return nil, err
	}

	// Create server
	s := Server{r: mux.NewRouter(), s: st, f: fn}

	return &s, nil
}

func (s *Server) Close() error {
	if s.f != nil {
		s.f()
	}
	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := start(r)
	rw := wrapper{http.StatusOK, w}

	// Panic recovery
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			stack := make([]byte, 1024*8)
			stack = stack[:runtime.Stack(stack, false)]

			tpl := "PANIC: %s\n%s"
			logger.Printf(tpl, err, stack)

			logger.Printf("panic: %+v", err)
			fmt.Fprintf(w, tpl, err, stack)
		}
	}()

	s.r.ServeHTTP(rw, r)

	t.end(logger, rw)
}

func (s *Server) success(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	// Set headers
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Prepare body
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// Send response
	w.WriteHeader(code)
	w.Write(b)
}

func (s *Server) failure(w http.ResponseWriter, r *http.Request, code int, cause error) {
	// Log error
	logger.Println(cause)

	// Set headers
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Prepare body
	data := map[string]string{"error": cause.Error()}
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// Send response
	w.WriteHeader(code)
	w.Write(b)
}

func (s *Server) parseBody(w http.ResponseWriter, r *http.Request, out interface{}) {
	d := json.NewDecoder(r.Body)
	if err := d.Decode(out); err != nil {
		if err != io.EOF {
			exc := fmt.Errorf("[parseBody] JSON parse error", err)
			s.failure(w, r, http.StatusBadRequest, exc)
			return
		}
	}
}
