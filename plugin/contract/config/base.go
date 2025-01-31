package config

type HTTPConfig struct {
	BaseUri  string `json:"baseUri,omitempty"`
	ProxyUri string `json:"proxyUri,omitempty"`
}

type PluginConfig struct {
	HTTPConfig
	AppId     string `json:"appId,omitempty"`
	AppSecret string `json:"appSecret,omitempty"`
	HttpDebug bool   `json:"httpDebug,omitempty"`
	Debug     bool   `json:"debug,omitempty"`
}
