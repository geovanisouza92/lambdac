package server

import (
	"fmt"
	"net/http"

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
	//
	// s.success(w, r, http.StatusCreated, functions)
}

func (s *Server) functionInfo(w http.ResponseWriter, r *http.Request) {
	//
	// s.success(w, r, http.StatusOK, functions)
}

func (s *Server) functionConfig(w http.ResponseWriter, r *http.Request) {
	//
	// s.success(w, r, http.StatusAccepted, functions)
}

func (s *Server) functionDestroy(w http.ResponseWriter, r *http.Request) {
	//
	// s.success(w, r, http.StatusGone, functions)
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
