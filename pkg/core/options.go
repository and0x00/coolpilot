package core

import (
	"flag"
)

type Options struct {
	Interactive bool
	Debug       bool
}

// ParseOptions parses the command line flags provided by a user
func ParseOptions() *Options {
	options := &Options{}
	flag.Parse()
	print(Banner())

	return options
}
