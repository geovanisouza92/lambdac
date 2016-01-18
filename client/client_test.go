package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	ErrInvalidURL = errors.New("Invalid URL")
)

func TestFunctionList(t *testing.T) {
	// TODO
}

func TestFunctionCreate(t *testing.T) {
	// TODO
}

func TestFunctionInfo(t *testing.T) {
	// TODO
}

func TestFunctionConfig(t *testing.T) {
	// TODO
}

func TestFunctionDestroy(t *testing.T) {
	// TODO
}

func TestFunctionEnv(t *testing.T) {
	// TODO
}

func TestFunctionEnvSet(t *testing.T) {
	// TODO
}

func TestFunctionEnvUnset(t *testing.T) {
	// TODO
}

func TestFunctionPull(t *testing.T) {
	// TODO
}

func TestFunctionPush(t *testing.T) {
	// TODO
}

func TestFunctionPs(t *testing.T) {
	// TODO
}

func TestFunctionLogs(t *testing.T) {
	// TODO
}

func TestFunctionStats(t *testing.T) {
	// TODO
}

func TestFunctionInvoke(t *testing.T) {
	// TODO
}

func TestRuntimeList(t *testing.T) {
	// TODO
}

func TestRuntimeCreate(t *testing.T) {
	// TODO
}

func TestRuntimeInfo(t *testing.T) {
	// TODO
}

func TestRuntimeDestroy(t *testing.T) {
	// TODO
}

func respondWith(t *testing.T, code int, path, body string) (*httptest.Server, API) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		fmt.Fprintln(w, body)
	}))

	tr := &http.Transport{
		Proxy: func(r *http.Request) (*url.URL, error) {
			if r.URL.Path != path {
				t.Errorf("Invalid path: expected %q got %q", path, r.URL.Path)
				return nil, ErrInvalidURL
			}
			return url.Parse(s.URL)
		},
	}

	hc := &http.Client{Transport: tr}

	c := New(s.URL, hc)

	return s, c
}
