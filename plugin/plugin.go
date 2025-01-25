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
		PluginName: "MediaXPlugin-Wechat",
	}
}

func (p *MediaXPlugin) Initialize(config map[string]interface{}) error {
	p.PluginName = "MediaXPlugin-Wechat"
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

var PluginMediaX MediaXPlugin = NewMediaXPlugin()
