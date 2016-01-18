package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type tracer struct {
	start time.Time
	data  map[string]interface{}
	sync.Mutex
}

func start(r *http.Request) (t *tracer) {
	t = &tracer{
		start: time.Now(),
		data: map[string]interface{}{
			"uri":     r.RequestURI,
			"method":  r.Method,
			"address": r.RemoteAddr,
		},
	}

	t.Lock()
	defer t.Unlock()

	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		t.data["address"] = ip
	}
	if id := r.Header.Get("X-Request-ID"); id != "" {
		t.data["id"] = id
	}

	return
}

func (t *tracer) end(l *log.Logger, w wrapper) {
	t.Lock()
	defer t.Unlock()

	t.data["elapsed"] = time.Since(t.start)
	t.data["status"] = w.status

	l.Println(t.String())
}

func (t *tracer) String() string {
	b := &bytes.Buffer{}

	for k, v := range t.data {
		fmt.Fprintf(b, "%v=%v ", k, v)
	}

	return b.String()
}
