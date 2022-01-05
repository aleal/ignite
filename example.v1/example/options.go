package example

import (
	"os"
)

// options for build the instance of any type you want.
type Options struct {
	Host    string `default:"localhost" desc:"example host"`
	Port    int    `default:"9999" desc:"example host"`
	Plugins struct {
		Custom struct {
			Enabled bool `default:"true" desc:"Custom plugin enabled."`
			Count   int  `config:"counterNumber" default:"490" desc:"Custom plugin enabled."`
		}
		Another struct {
			Enabled bool   `default:"true" desc:"another plugin enabled."`
			Label   string `default:"LabelValue123" desc:"another plugin label."`
		}
	}
}

// root path for the options
func (o *Options) Root() string {
	return "ignite.example"
}

// post load options
func (o *Options) PostLoad() error {
	// check env for overriding options (for example)
	if host := os.Getenv("EXAMPLE_HOST"); host != "" {
		o.Host = host
	}
	return nil
}
