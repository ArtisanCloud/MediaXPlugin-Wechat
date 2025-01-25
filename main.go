package main

import (
	"fmt"
	plugin2 "github.com/ArtisanCloud/MediaXCore/pkg/plugin"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	"plugin"
)

func main() {
	// 加载插件
	p, err := plugin.Open("./plugin/plugin.so")
	if err != nil {
		panic(err)
	}

	// 加载yaml配置文件
	configPlugin, err := plugin2.ReadPluginMetadata("./plugin/plugin.yaml")
	if err != nil {
		panic(err)
	}
	pluginName := configPlugin.Name

	// 获取插件对象
	mediaXPlugin, err := plugin2.LookUpSymbol[contract.ProviderInterface](p, pluginName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("plugin loaded name :%s", (*mediaXPlugin).Name())

}
