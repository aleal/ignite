package plugins

import (
	"context"

	"github.com/aleal/ignite/example.v1/example"
)

// another custom plugin
func AnotherCustom(ctx context.Context, w *example.Wrapper) error {
	opts := w.Options.Plugins.Another
	// checks if plugin is enabled
	if opts.Enabled {
		// plugin code goes here
		w.Instance.Label = opts.Label
	}
	return nil
}
