package example

import (
	"os"
	"reflect"
	"testing"

	"github.com/aleal/ignite"
)

func TestOptions(t *testing.T) {
	tests := []struct {
		name  string
		setup func()
		want  func() *Options
	}{
		{
			name:  "loads options",
			setup: func() {},
			want: func() *Options {
				opts := ignite.New[*Options]()
				opts.Host = "localhost"
				opts.Port = 9999
				opts.Plugins.Custom.Enabled = true
				opts.Plugins.Custom.Count = 490
				opts.Plugins.Another.Enabled = true
				opts.Plugins.Another.Label = "LabelValue123"
				return opts
			},
		},
		{
			name: "loads options by overriding host for an env var",
			setup: func() {
				os.Setenv("EXAMPLE_HOST", "example.com")
			},
			want: func() *Options {
				opts := ignite.New[*Options]()
				opts.Host = "example.com"
				opts.Port = 9999
				opts.Plugins.Custom.Enabled = true
				opts.Plugins.Custom.Count = 490
				opts.Plugins.Another.Enabled = true
				opts.Plugins.Another.Label = "LabelValue123"
				return opts
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, e := ignite.Load[*Options]()
			if e != nil {
				t.Errorf("Unexpected error %v", e)
			}
			want := tt.want()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("want\t%v\ngot \t%v", want, got)
			}

		})
	}
}
