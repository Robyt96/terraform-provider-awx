package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/robyt96/terraform-provider-awx/awx"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return awx.Provider()
		},
	}

	if debug {
		opts.Debug = true
		opts.ProviderAddr = "terraform.local/local/awx"
	}

	plugin.Serve(opts)
}
