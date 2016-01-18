package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/geovanisouza92/env"
	"github.com/geovanisouza92/lambdac/types"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func (s *Server) functionList(w http.ResponseWriter, r *http.Request) {
	// TODO Include filters

	functions, err := s.s.Functions().All()
	if err != nil {
		exc := fmt.Errorf("[functionList] Store error caused by: %s", err)
		s.failure(w, r, http.StatusInternalServerError, exc)
		return
	}

	s.success(w, r, http.StatusOK, functions)
}

func (s *Server) functionCreate(w http.ResponseWriter, r *http.Request) {
	var f types.Function
	s.parseBody(w, r, &f)

	// Validate function
	if f.Name == "" {
		exc := fmt.Errorf("[functionCreate] Attribute \"name\" is required")
		s.failure(w, r, http.StatusBadRequest, exc)
		return
	}
	if f.Runtime == "" {
		exc := fmt.Errorf("[functionCreate] Attribute \"runtime\" is required")
		s.failure(w, r, http.StatusBadRequest, exc)
		return
	}
	if f.Timeout <= 0 {
		exc := fmt.Errorf("[functionCreate] Attribute \"timeout\" must be greater than 0")
		s.failure(w, r, http.StatusBadRequest, exc)
		return
	}
	if f.Memory < 32 {
		exc := fmt.Errorf("[functionCreate] Attribute \"memory\" must be greater than 32 MB")
		s.failure(w, r, http.StatusBadRequest, exc)
		return
	}
	if f.Instances <= 0 {
		exc := fmt.Errorf("[functionCreate] Attribute \"instances\" must be greater than 0")
		s.failure(w, r, http.StatusBadRequest, exc)
		return
	}
	for _, e := range f.Env {
		if _, _, err := env.ParseLine(e); err != nil {
			exc := fmt.Errorf("Invalid environment variable %q. Pattern must be VARIABLE=VALUE", e)
			s.failure(w, r, http.StatusBadRequest, exc)
			return
		}
	}

	now := time.Now()
	f.ID = types.NewID()
	f.Created = now
	f.Updated = now

	created, err := s.s.Functions().Create(f)
	if err != nil {
		exc := fmt.Errorf("[functionCreate] Store error caused by: %s", err)
		s.failure(w, r, http.StatusInternalServerError, exc)
		return
	}
	logger.Printf("Function %q created with runtime %q", created.Name, created.Runtime)

	s.success(w, r, http.StatusCreated, created)
}

func (s *Server) functionInfo(w http.ResponseWriter, r *http.Request) {
	function := context.Get(r, "function").(types.Function)
	s.success(w, r, http.StatusOK, function)
}

func (s *Server) functionConfig(w http.ResponseWriter, r *http.Request) {
	//
	// s.success(w, r, http.StatusAccepted, functions)
}

func (s *Server) functionDestroy(w http.ResponseWriter, r *http.Request) {
	function := context.Get(r, "function").(types.Function)

	s.functionDestroyInternal(w, r, function)

	s.success(w, r, http.StatusGone, nil)
}

func (s *Server) functionDestroyInternal(w http.ResponseWriter, r *http.Request, function types.Function) {
	// TODO Destroy function instances

	// Destroy function
	err := s.s.Functions().Remove(function.ID)
	if err != nil {
		exc := fmt.Errorf("[functionDestroy] Store error caused by: %s\n", err)
		s.failure(w, r, http.StatusInternalServerError, exc)
		return
	}
}

func (s *Server) functionEnv(w http.ResponseWriter, r *http.Request) {
	//
	// s.success(w, r, http.StatusOK, functions)
}

func (s *Server) functionEnvSet(w http.ResponseWriter, r *http.Request) {
	//
	// s.success(w, r, http.StatusAccepted, functions)
}

func (s *Server) functionEnvUnset(w http.ResponseWriter, r *http.Request) {
	//
	// s.success(w, r, http.StatusGone, functions)
}

func (s *Server) functionPull(w http.ResponseWriter, r *http.Request) {
	//
	// s.success(w, r, http.StatusOK, functions)
}

func (s *Server) functionPush(w http.ResponseWriter, r *http.Request) {
	//
	// s.success(w, r, http.StatusAccepted, functions)
}

func (s *Server) functionPs(w http.ResponseWriter, r *http.Request) {
	//
	// s.success(w, r, http.StatusOK, functions)
}

func (s *Server) functionLogs(w http.ResponseWriter, r *http.Request) {
	//
	// s.success(w, r, http.StatusOK, functions)
}

func (s *Server) functionStats(w http.ResponseWriter, r *http.Request) {
	//
	// s.success(w, r, http.StatusOK, functions)
}

func (s *Server) functionInvoke(w http.ResponseWriter, r *http.Request) {
	//
	// s.success(w, r, http.StatusAccepted, functions)
}

func (s *Server) queryFunction(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		// Find the function
		function, err := s.s.Functions().FindByIDOrName(id)
		if err != nil {
			exc := fmt.Errorf("[queryFunction] Store error caused by: %s", err)
			s.failure(w, r, http.StatusInternalServerError, exc)
			return
		}

		// Put function information in context
		context.Set(r, "function", function)

		next(w, r)
	}
}
