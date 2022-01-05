package example

import (
	"context"
	"reflect"
	"testing"

	"github.com/aleal/ignite"
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
				opts.Host = "localhost"
				opts.Port = 9999
				opts.Plugins.Custom.Enabled = true
				opts.Plugins.Custom.Count = 490
				opts.Plugins.Another.Enabled = true
				opts.Plugins.Another.Label = "LabelValue123"
				return opts
			},
			want: func() *Any {
				return &Any{
					Host: "localhost",
					Port: 9999,
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
