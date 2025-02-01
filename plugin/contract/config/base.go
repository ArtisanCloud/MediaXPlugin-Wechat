package config

import "github.com/ArtisanCloud/MediaXCore/pkg/logger/config"

type HTTPConfig struct {
	BaseUri  string `json:"baseUri,omitempty"`
	ProxyUri string `json:"proxyUri,omitempty"`
}

type PluginConfig struct {
	HTTPConfig
	AppId     string `json:"appId,omitempty"`
	AppSecret string `json:"appSecret,omitempty"`
	LogConfig config.LogConfig
}
