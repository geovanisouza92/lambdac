// Types used on server and client
package types

import (
	"time"
)

type Function struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Runtime     string    `json:"runtime,omitempty"`
	Handler     string    `json:"handler,omitempty"`
	Description string    `json:"description,omitempty"`
	Code        string    `json:"code,omitempty"`
	Timeout     int64     `json:"timeout,omitempty"`
	Memory      int       `json:"memory,omitempty"`
	Instances   int       `json:"instances,omitempty"`
	Env         []string  `json:"env,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	Updated     time.Time `json:"updated,omitempty"`
}

type Functions []Function

type Runtime struct {
	ID      string    `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Label   string    `json:"label,omitempty"`
	Image   string    `json:"image,omitempty"`
	Command string    `json:"command,omitempty"`
	Agent   bool      `json:"agent,omitempty"`
	Driver  string    `json:"driver,omitempty"`
	Options []string  `json:"options,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
}

type Runtimes []Runtime
