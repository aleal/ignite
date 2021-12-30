package plugins

import (
	"github.com/aleal/ignite"
	"github.com/aleal/ignite/example.v1/example"
)

var All = []ignite.Plugin[*example.Wrapper, *example.Options]{
	Custom,
	AnotherCustom,
}
