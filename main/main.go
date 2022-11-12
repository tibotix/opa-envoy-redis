package main

import (
	"os"

	"github.com/open-policy-agent/opa/cmd"
	"github.com/open-policy-agent/opa/runtime"
	opa_redis_plugin "github.com/tibotix/opa-redis-plugin/plugin"
	opa_envoy_plugin "github.com/open-policy-agent/opa-envoy-plugin/plugin"
)

func main() {
	runtime.RegisterPlugin(opa_redis_plugin.PluginName, opa_redis_plugin.Factory{})
	runtime.RegisterPlugin(opa_envoy_plugin.PluginName, opa_envoy_plugin.Factory{})

	if err := cmd.RootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
