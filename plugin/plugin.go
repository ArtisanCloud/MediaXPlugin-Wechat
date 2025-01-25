package main

import (
	"fmt"

	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
)

type MediaXPlugin struct {
	PluginName string
}

func NewMediaXPlugin() MediaXPlugin {
	return MediaXPlugin{
		PluginName: "PluginMediaXWechat",
	}
}

func (p *MediaXPlugin) Initialize(config map[string]interface{}) error {
	p.PluginName = "PluginMediaXWechat"
	return nil
}

func (p *MediaXPlugin) Name() string {
	return p.PluginName
}

func (p *MediaXPlugin) Publish(*contract.PublishRequest, ...interface{}) (*contract.PublishResult, error) {

	fmt.Println("Publishing MediaX Plugin")
	return &contract.PublishResult{
		Status:  "success",
		Message: "MediaX Wechat Plugin Published Successfully",
	}, nil
}

var PluginMediaXWechat MediaXPlugin = NewMediaXPlugin()
