package server

import (
	"net/http"
)

type wrapper struct {
	status int
	http.ResponseWriter
}

func (w wrapper) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}
