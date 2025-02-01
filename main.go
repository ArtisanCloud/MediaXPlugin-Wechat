package main

import (
	"context"
	"github.com/ArtisanCloud/MediaXCore/pkg/logger"
	"github.com/ArtisanCloud/MediaXCore/pkg/logger/config"
	"plugin"

	plugin2 "github.com/ArtisanCloud/MediaXCore/pkg/plugin"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
)

func main() {

	ctx := context.Background()
	configPlugin := &contract.PluginConfig{
		LogConfig: config.LogConfig{
			Level:         "debug",
			Console:       true,
			UseJsonFormat: true,
			File: config.FileConfig{
				Enable: true,
				//InfoFilePath:  "./logs/info.log",
				//ErrorFilePath: "./logs/error.log",
			},
			//Loki: config.LokiConfig{},
			HttpDebug: true,
			Debug:     true,
		},
	}

	log := logger.GetLogger(&configPlugin.LogConfig)

	// 加载yaml配置文件
	configPluginsMetadata, err := plugin2.ReadPluginMetadata("./plugins/plugins.yaml")
	if err != nil {
		panic(err)
	}

	log.Info("parsed plugin bundle config file " + configPluginsMetadata.Name)

	for _, configPluginMetadata := range configPluginsMetadata.Plugins {

		pluginName := configPluginMetadata.Name
		buildPath := configPluginMetadata.BuildPath

		// 加载插件hgnba        pki0ujhbvg cvi987t6re34578765o766789
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

		// 初始化plugin
		err = mediaXPlugin.Initialize(&ctx, configPlugin)
		if err != nil {
			panic(err)
		}

		// 获取插件名称
		log.InfoF("plugin loaded name :%s", mediaXPlugin.Name(&ctx))

		// 插件发布内容
		resObj, err := mediaXPlugin.Publish(&ctx, &contract.PublishRequest{
			Content: "Hello, MediaX Plugin!",
		})

		if err != nil {
			panic(err)
		}
		if obj := resObj.(*contract.PublishResponse); obj.Code == 0 {
			log.Info("Publishing MediaX Plugin")
		} else {
			log.Info("Failed to publish MediaX Plugin")
		}
	}

}
