package example

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/aleal/ignite"
	"github.com/aleal/ignite/example.v1/example"
	"github.com/americanas-go/config"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		setup func()
		want  func() *example.Any
	}{
		{
			name: "new instance with default options",
			setup: func() {
				config.Load()
			},
			want: func() *example.Any {
				return &example.Any{
					Host:  "localhost",
					Port:  9999,
					Count: 70 * 7,
					Label: "Label",
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, e := New(context.Background())
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

func TestNewWithConfigPath(t *testing.T) {
	tests := []struct {
		name  string
		setup func()
		want  func() *example.Any
	}{
		{
			name: "new instance with config path",
			setup: func() {
				os.Setenv("CUSTOM_HOST", "192.68.1.1")
				os.Setenv("CUSTOM_PORT", "7777")
				os.Setenv("CUSTOM_PLUGINS_CUSTOM_ENABLED", "true")
				os.Setenv("CUSTOM_PLUGINS_CUSTOM_COUNT", "18")
				os.Setenv("CUSTOM_PLUGINS_ANOTHER__CUSTOM_ENABLED", "true")
				os.Setenv("CUSTOM_PLUGINS_ANOTHER__CUSTOM_LABEL", "ANOTHER_LABEL")
				config.Load()
			},
			want: func() *example.Any {
				return &example.Any{
					Host:  "192.68.1.1",
					Port:  7777,
					Count: 18,
					Label: "ANOTHER_LABEL",
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, e := NewWithConfigPath(context.Background(), "custom")
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

func TestNewWithOptions(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *example.Options
		want  func() *example.Any
	}{
		{
			name: "new instance with options",
			setup: func() *example.Options {
				opts := ignite.New[*example.Options]()
				opts.Host = "example.com"
				opts.Port = 8888
				opts.Plugins.Custom.Enabled = false // custom plugin disabled
				opts.Plugins.Custom.Count = 4444
				opts.Plugins.AnotherCustom.Enabled = true
				opts.Plugins.AnotherCustom.Label = "BLAH"
				return opts
			},
			want: func() *example.Any {
				return &example.Any{
					Host:  "example.com",
					Port:  8888,
					Count: 0, // zero value
					Label: "BLAH",
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := tt.setup()
			got, e := NewWithOptions(context.Background(), opts)
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
