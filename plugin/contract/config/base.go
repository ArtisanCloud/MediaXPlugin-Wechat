package config

type HTTPConfig struct {
	BaseUri  string `json:"baseUri,omitempty"`
	ProxyUri string `json:"proxyUri,omitempty"`
}

type PluginConfig struct {
	HTTPConfig
	AppId     string `json:"appId,omitempty"`
	AppScret  string `json:"appScret,omitempty"`
	HttpDebug bool   `json:"httpDebug,omitempty"`
	Debug     bool   `json:"debug,omitempty"`
}
