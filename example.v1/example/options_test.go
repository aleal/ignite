package example

import (
	"os"
	"reflect"
	"testing"

	"github.com/aleal/ignite"
	"github.com/americanas-go/config"
)

func TestOptions(t *testing.T) {
	tests := []struct {
		name  string
		setup func()
		want  func() *Options
	}{
		{
			name: "loads options",
			setup: func() {
				AddConfig(root)
				config.Load()
			},
			want: func() *Options {
				opts := ignite.New[*Options]()
				opts.Host = config.String(root + host)
				opts.Port = config.Int(root + port)
				opts.Plugins.Custom.Enabled = config.Bool(root + pluginCustomEnabled)
				opts.Plugins.Custom.Count = config.Int(root + pluginCustomCount)
				opts.Plugins.AnotherCustom.Enabled = config.Bool(root + pluginAnotherCustomEnabled)
				opts.Plugins.AnotherCustom.Label = config.String(root + pluginAnotherCustomLabel)
				return opts
			},
		},
		{
			name: "loads options by overriding host for an env var",
			setup: func() {
				AddConfig(root)
				os.Setenv("EXAMPLE_HOST", "example.com")
				config.Load()
			},
			want: func() *Options {
				opts := ignite.New[*Options]()
				opts.Host = "example.com"
				opts.Port = config.Int(root + port)
				opts.Plugins.Custom.Enabled = config.Bool(root + pluginCustomEnabled)
				opts.Plugins.Custom.Count = config.Int(root + pluginCustomCount)
				opts.Plugins.AnotherCustom.Enabled = config.Bool(root + pluginAnotherCustomEnabled)
				opts.Plugins.AnotherCustom.Label = config.String(root + pluginAnotherCustomLabel)
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
