package main

import (
	"context"
	"fmt"
	"plugin"

	plugin2 "github.com/ArtisanCloud/MediaXCore/pkg/plugin"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	contract2 "github.com/ArtisanCloud/MediaXPlugin/plugin/contract"
)

func main() {
	// 加载yaml配置文件
	configPlugin, err := plugin2.ReadPluginMetadata("./plugin/plugin.yaml")
	if err != nil {
		panic(err)
	}
	pluginName := configPlugin.Name
	buildPath := configPlugin.BuildPath

	// 加载插件
	p, err := plugin.Open(buildPath)
	if err != nil {
		panic(err)
	}

	// 获取插件对象
	obj, err := plugin2.LookUpSymbol[contract.ProviderInterface](p, pluginName)
	if err != nil {
		panic(err)
	}
	mediaXPlugin := *obj

	// 获取插件名称
	ctx := context.Background()
	fmt.Printf("plugin loaded name :%s \n", mediaXPlugin.Name(&ctx))

	// 插件发布内容
	resObj, err := mediaXPlugin.Publish(&ctx, &contract2.PublishRequest{
		Content: "Hello, MediaX Plugin!",
	})

	if err != nil {
		panic(err)
	}
	if obj := resObj.(*contract2.PublishResponse); obj.Code == 0 {
		fmt.Println("Publishing MediaX Plugin")
	} else {
		fmt.Println("Failed to publish MediaX Plugin")
	}

}
