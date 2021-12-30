package example

import (
	"os"

	"github.com/americanas-go/config"
)

// options for build the instance of any type you want.
type Options struct {
	Host    string
	Port    int
	Plugins struct {
		Custom struct {
			Enabled bool
			Count   int
		}
		AnotherCustom struct {
			Enabled bool
			Label   string
		}
	}
}

var optionsRoot = ""

// root path for the options
func (o *Options) Root() string {
	if optionsRoot != "" {
		return optionsRoot
	}
	return root
}

// post load options
func (o *Options) PostLoad() error {
	// check env for overriding options (for example)
	if host := os.Getenv("EXAMPLE_HOST"); host != "" {
		o.Host = host
	}
	return nil
}

// options configuration
const (
	root                       = "ignite.example"
	host                       = ".host"
	port                       = ".port"
	pluginsRoot                = ".plugins"
	pluginCustomRoot           = pluginsRoot + ".custom"
	pluginCustomEnabled        = pluginCustomRoot + ".enabled"
	pluginCustomCount          = pluginCustomRoot + ".count"
	pluginAnotherCustomRoot    = pluginsRoot + ".anotherCustom"
	pluginAnotherCustomEnabled = pluginAnotherCustomRoot + ".enabled"
	pluginAnotherCustomLabel   = pluginAnotherCustomRoot + ".label"
)

func init() {
	AddConfig(root)
}

// adds default config values at path.
func AddConfig(path string) {
	optionsRoot = path
	// root options
	config.Add(path+host, "localhost", "client host")
	config.Add(path+port, 9999, "client port")

	// plugins options if any
	config.Add(path+pluginCustomEnabled, true, "custom plugin enabled")
	config.Add(path+pluginCustomCount, 70*7, "custom plugin forgiveness count number")
	config.Add(path+pluginAnotherCustomEnabled, true, "another custom plugin enabled")
	config.Add(path+pluginAnotherCustomLabel, "Label", "another custom plugin label")
}
