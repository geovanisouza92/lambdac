package server

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/geovanisouza92/lambdac/driver"
	_ "github.com/geovanisouza92/lambdac/driver/docker"
	"github.com/geovanisouza92/lambdac/types"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

var nonLetter, _ = regexp.Compile("[^a-zA-Z0-9]")

func (s *Server) runtimeList(w http.ResponseWriter, r *http.Request) {
	data, err := s.s.Runtimes().All()
	if err != nil {
		exc := fmt.Errorf("[runtimeList] Store error caused by: %s", err)
		s.failure(w, r, http.StatusInternalServerError, exc)
		return
	}

	s.success(w, r, http.StatusOK, data)
}

func (s *Server) runtimeCreate(w http.ResponseWriter, r *http.Request) {
	var rt types.Runtime
	s.parseBody(w, r, &rt)

	// Validate runtime
	if rt.Name == "" {
		exc := fmt.Errorf("[runtimeCreate] Attribute \"name\" is required")
		s.failure(w, r, http.StatusBadRequest, exc)
		return
	}
	if rt.Image == "" {
		exc := fmt.Errorf("[runtimeCreate] Attribute \"image\" is required")
		s.failure(w, r, http.StatusBadRequest, exc)
		return
	}
	if rt.Command == "" {
		exc := fmt.Errorf("[runtimeCreate] Attribute \"command\" is required")
		s.failure(w, r, http.StatusBadRequest, exc)
		return
	}
	if rt.Driver == "" {
		exc := fmt.Errorf("[runtimeCreate] Attribute \"driver\" is required")
		s.failure(w, r, http.StatusBadRequest, exc)
		return
	}
	_, err := driver.Open(rt.Driver)
	if err != nil {
		exc := fmt.Errorf("[runtimeCreate] Invalid driver: %s", err)
		s.failure(w, r, http.StatusBadRequest, exc)
		return
	}

	// Set other values
	now := time.Now()
	rt.ID = types.NewID()
	rt.Created = now
	rt.Updated = now

	// Save runtime configuration
	created, err := s.s.Runtimes().Create(rt)
	if err != nil {
		exc := fmt.Errorf("[runtimeCreate] Store error caused by: %s", err)
		s.failure(w, r, http.StatusInternalServerError, exc)
		return
	}
	logger.Printf("Runtime %q created with image %q and command %q.\n", created.Name, created.Image, created.Command)

	s.success(w, r, http.StatusCreated, created)
}

func (s *Server) runtimeInfo(w http.ResponseWriter, r *http.Request) {
	runtime := context.Get(r, "runtime").(types.Runtime)
	s.success(w, r, http.StatusOK, runtime)
}

func (s *Server) runtimeDestroy(w http.ResponseWriter, r *http.Request) {

	// Get runtime param
	runtime := context.Get(r, "runtime").(types.Runtime)

	// Load related functions
	functions, err := s.s.Functions().FindByRuntimeID(runtime.ID)
	if err != nil {
		exc := fmt.Errorf("[runtimeDestroy] Store error caused by: %s", err)
		s.failure(w, r, http.StatusInternalServerError, exc)
		return
	}

	// Load options
	q := r.URL.Query()
	force, err := strconv.ParseBool(q.Get("force"))
	if err != nil {
		exc := fmt.Errorf("[runtimeDestroy] Invalid \"force\" value: %s", err)
		s.failure(w, r, http.StatusBadRequest, exc)
		return
	}

	// If not --force, say to user destroy related functions first
	if len(functions) > 0 && !force {
		exc := fmt.Errorf("[runtimeDestroy] This runtime is used by other functions. Destroy them first or use --force option")
		s.failure(w, r, http.StatusBadRequest, exc)
		return
	}

	// Destroy related functions
	for _, f := range functions {
		s.functionDestroyInternal(w, r, f)
	}

	// Destroy runtime
	err = s.s.Runtimes().Remove(runtime.ID)
	if err != nil {
		exc := fmt.Errorf("[runtimeDestroy] Store error caused by: %s\n", err)
		s.failure(w, r, http.StatusInternalServerError, exc)
		return
	}

	s.success(w, r, http.StatusGone, nil)
}

func (s *Server) queryRuntime(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		// Find the runtime
		runtime, err := s.s.Runtimes().FindByIDOrName(id)
		if err != nil {
			exc := fmt.Errorf("[queryRuntime] Store error caused by: %s", err)
			s.failure(w, r, http.StatusInternalServerError, exc)
			return
		}

		// Put runtime information in context
		context.Set(r, "runtime", runtime)

		next(w, r)
	}
}
