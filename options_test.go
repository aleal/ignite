package ignite

import (
	"os"
	"reflect"
	"testing"

	"github.com/americanas-go/config"
)

type CustomOptions struct {
	Enabled             bool
	Host                string
	Port                int
	CustomPluginOptions *CustomPluginOptions
}

func (co *CustomOptions) Root() string {
	return "ignite.custom"
}

func (co *CustomOptions) PostLoad() (e error) {
	co.CustomPluginOptions, e = Load[*CustomPluginOptions]()
	return
}

type CustomPluginOptions struct {
	Enabled bool
	Count   int
}

func (co *CustomPluginOptions) Root() string {
	return "ignite.custom.plugins.custom"
}

func (co *CustomPluginOptions) PostLoad() error {
	return nil
}

func TestLoadOptions(t *testing.T) {
	os.Setenv("IGNITE_CUSTOM_ENABLED", "true")
	os.Setenv("IGNITE_CUSTOM_HOST", "localhost")
	os.Setenv("IGNITE_CUSTOM_PORT", "9999")
	os.Setenv("IGNITE_CUSTOM_PLUGINS_CUSTOM_ENABLED", "true")
	os.Setenv("IGNITE_CUSTOM_PLUGINS_CUSTOM_COUNT", "18")
	config.Load()
	tests := []struct {
		name    string
		want    *CustomOptions
		wantErr func(error) bool
	}{
		{
			name: "Returns new ignite options",
			want: &CustomOptions{
				Enabled: true,
				Host:    "localhost",
				Port:    9999,
				CustomPluginOptions: &CustomPluginOptions{
					Enabled: true,
					Count:   18,
				},
			},
			wantErr: func(e error) bool { return e == nil },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, e := Load[*CustomOptions]()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want\t%v\ngot \t%v", tt.want, got)
			}
			if !tt.wantErr(e) {
				t.Errorf("Unexpected error %v", e)
			}
		})
	}
}
