package example

import (
	"context"

	"github.com/aleal/ignite"
	"github.com/aleal/ignite/example.v1/example"
	"github.com/aleal/ignite/example.v1/plugins"
)

// creates a new instance of the desired type with default options.
func New(ctx context.Context) (*example.Any, error) {
	w, e := ignite.Setup(ctx, plugins.All...)
	if e != nil {
		return nil, e
	}
	return w.Instance, nil
}

// creates a new instance of the desired type with options from config path.
func NewWithConfigPath(ctx context.Context, path string) (*example.Any, error) {
	w, e := ignite.SetupWithConfigPath(ctx, path, plugins.All...)
	if e != nil {
		return nil, e
	}
	return w.Instance, nil
}

// creates a new instance of the desired type with options.
func NewWithOptions(ctx context.Context, o *example.Options) (*example.Any, error) {
	w, e := ignite.SetupWithOptions(ctx, o, plugins.All...)
	if e != nil {
		return nil, e
	}
	return w.Instance, nil
}
