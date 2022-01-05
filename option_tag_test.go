package ignite

import (
	"reflect"
	"testing"
)

type TestOptionTagsOptions struct {
	Host    string `default:"localhost" desc:"Http server host."`
	Port    int    `default:"9999"      desc:"Http server port."`
	Plugins struct {
		Custom struct {
			Enabled bool `config:"blah" default:"false" desc:"whether custom plugin is enabled"`
		} `desc:"custom plugin options"`
		Another struct {
			Enabled bool `default:"false" desc:"whether another plugin is enabled"`
		} `desc:"another plugin options"`
	} `desc:"plugins options"`
}

func (o TestOptionTagsOptions) Root() string {
	return "ignite.test"
}

func (o TestOptionTagsOptions) PostLoad() error {
	return nil
}

func TestGetTags(t *testing.T) {
	tests := []struct {
		name string
		opts IgniteOptions
		want func() []*IgniteOptionTag
	}{
		{
			name: "Returns ignite option tags from options struct pointer",
			opts: &TestOptionTagsOptions{},
			want: func() []*IgniteOptionTag {
				return []*IgniteOptionTag{
					{
						Config:      "host",
						Default:     "localhost",
						Description: "Http server host.",
						Path:        ".host",
					},
					{
						Config:      "port",
						Default:     "9999",
						Description: "Http server port.",
						Path:        ".port",
					},
					{
						Config:      "blah",
						Default:     "false",
						Description: "whether custom plugin is enabled",
						Path:        ".plugins.custom.blah",
					},
					{
						Config:      "enabled",
						Default:     "false",
						Description: "whether another plugin is enabled",
						Path:        ".plugins.another.enabled",
					},
				}
			},
		},
		{
			name: "Returns ignite option tags from options struct",
			opts: TestOptionTagsOptions{},
			want: func() []*IgniteOptionTag {
				return []*IgniteOptionTag{
					{
						Config:      "host",
						Default:     "localhost",
						Description: "Http server host.",
						Path:        ".host",
					},
					{
						Config:      "port",
						Default:     "9999",
						Description: "Http server port.",
						Path:        ".port",
					},
					{
						Config:      "blah",
						Default:     "false",
						Description: "whether custom plugin is enabled",
						Path:        ".plugins.custom.blah",
					},
					{
						Config:      "enabled",
						Default:     "false",
						Description: "whether another plugin is enabled",
						Path:        ".plugins.another.enabled",
					},
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTags(tt.opts)
			want := tt.want()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("\nwant\t%v\ngot \t%v", want, got)
			}
		})
	}
}
