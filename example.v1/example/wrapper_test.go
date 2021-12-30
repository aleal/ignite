package example

import (
	"context"
	"reflect"
	"testing"

	"github.com/aleal/ignite"
	"github.com/americanas-go/config"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *Options
		want  func() *Any
	}{
		{
			name: "initialize instance",
			setup: func() *Options {
				opts := ignite.New[*Options]()
				opts.Host = config.String(root + host)
				opts.Port = config.Int(root + port)
				opts.Plugins.Custom.Enabled = config.Bool(root + pluginCustomEnabled)
				opts.Plugins.Custom.Count = config.Int(root + pluginCustomCount)
				opts.Plugins.AnotherCustom.Enabled = config.Bool(root + pluginAnotherCustomEnabled)
				opts.Plugins.AnotherCustom.Label = config.String(root + pluginAnotherCustomLabel)
				return opts
			},
			want: func() *Any {
				return &Any{
					Host: config.String(root + host),
					Port: config.Int(root + port),
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := tt.setup()
			wrapper := ignite.New[*Wrapper]()
			e := wrapper.Init(context.Background(), opts)
			if e != nil {
				t.Errorf("Unexpected error %v", e)
			}
			want := tt.want()
			got := wrapper.Instance
			if !reflect.DeepEqual(got, want) {
				t.Errorf("want\t%v\ngot \t%v", want, got)
			}

		})
	}
}
