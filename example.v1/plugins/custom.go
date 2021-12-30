package plugins

import (
	"context"

	"github.com/aleal/ignite/example.v1/example"
)

// a custom plugin
func Custom(ctx context.Context, w *example.Wrapper) error {
	opts := w.Options.Plugins.Custom
	// checks if plugin is enabled
	if opts.Enabled {
		// plugin code goes here
		w.Instance.Count = opts.Count
	}
	return nil
}
