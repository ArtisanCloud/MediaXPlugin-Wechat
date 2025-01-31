package main

import (
	"context"
	"fmt"
	"reflect"

	contract2 "github.com/ArtisanCloud/MediaXPlugin/plugin/contract"
	"github.com/ArtisanCloud/MediaXPlugin/plugin/contract/config"
)

type MediaXPlugin struct {
	PluginName string
}

func NewMediaXPlugin() MediaXPlugin {
	return MediaXPlugin{
		PluginName: "PluginMediaX",
	}
}

func (p *MediaXPlugin) Initialize(ctx *context.Context, arg interface{}) error {

	// parse arg to contract config.PluginConfig firstly
	c, ok := arg.(*config.PluginConfig)
	if !ok {
		argType := reflect.TypeOf(arg)
		return fmt.Errorf("invalid argument type %s for PluginMediaX arg: *config.PluginConfig", argType.String())
	}

	p.PluginName = "PluginMediaX"
	fmt.Printf("Initializing %s plugin with config base uri: %+s\n", p.PluginName, c.BaseUri)

	return nil
}

func (p *MediaXPlugin) Name(ctx *context.Context) string {
	return p.PluginName
}

func (p *MediaXPlugin) Publish(ctx *context.Context, arg interface{}) (interface{}, error) {

	// parse arg to contract contract2.PublishRequest firstly
	req, ok := arg.(*contract2.PublishRequest)
	if !ok {
		argType := reflect.TypeOf(arg)
		return nil, fmt.Errorf("invalid argument type %s for PluginMediaX arg: *contract.PublishRequest", argType.String())
	}
	fmt.Printf("Publishing %s plugin with request: %+s\n", p.PluginName, req.Content)

	// prepare response data
	result := &contract2.PublishResponse{}
	result.Code = 0
	result.Msg = "MediaX Plugin Published Successfully"

	return result, nil
}

var PluginMediaX MediaXPlugin = NewMediaXPlugin()
