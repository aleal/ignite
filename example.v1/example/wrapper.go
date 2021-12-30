package example

import (
	"context"
)

// any type you want to build.
// A server, a client, a connection and so on.
type Any struct {
	Host  string
	Port  int
	Count int
	Label string
}

// ignite instance wrapping your desired instance of any type.
type Wrapper struct {
	Instance *Any
	Options  *Options
}

// initiates building your object with configured options.
func (w *Wrapper) Init(ctx context.Context, o *Options) error {
	w.Options = o
	w.Instance = &Any{
		Host: o.Host,
		Port: o.Port,
	}
	return nil
}
