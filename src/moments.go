package main

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/MediaXCore/pkg/logger"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	"reflect"
)

type MediaXPluginWechatMoment struct {
	PluginName string
	Logger     *logger.Logger
}

func NewMediaXPluginWechatMoment() MediaXPluginWechatMoment {
	return MediaXPluginWechatMoment{
		PluginName: string(contract.WechatMoments),
	}
}

func (p *MediaXPluginWechatMoment) Initialize(ctx *context.Context, arg interface{}) error {

	// parse arg to contract config.PluginConfig firstly
	c, ok := arg.(*contract.PluginConfig)
	if !ok {
		argType := reflect.TypeOf(arg)
		return fmt.Errorf("initializing %s invalid argument type %s for arg: *config.PluginConfig", p.PluginName, argType.String())
	}
	// 初始化logger
	p.Logger = logger.NewLogger(&c.LogConfig)

	p.Logger.InfoF("Initializing %s plugin with config base uri: %+s\n", p.PluginName, c.BaseUri)

	return nil
}

func (p *MediaXPluginWechatMoment) Name(ctx *context.Context) string {
	return p.PluginName
}

func (p *MediaXPluginWechatMoment) Publish(ctx *context.Context, arg interface{}) (interface{}, error) {

	// parse arg to contract contract2.PublishRequest firstly
	req, ok := arg.(*contract.PublishRequest)
	if !ok {
		argType := reflect.TypeOf(arg)
		return nil, fmt.Errorf("invalid argument type %s for PluginMediaX arg: *contract.PublishRequest", argType.String())
	}
	p.Logger.InfoF("Publishing %s plugin with request: %+s\n", p.PluginName, req.Content)

	// do what you like here with the request and return a response
	// ...

	// prepare response data
	result := &contract.PublishResponse{}
	result.Code = 0
	result.Msg = "MediaX Plugin Published Successfully"

	return result, nil
}

// make sure the variable is exported as the same "Name" as the name in the plugin.yaml file.
var WechatMoments MediaXPluginWechatMoment = NewMediaXPluginWechatMoment()
