# MediaXPlugin-XXX 组件

本仓库基于 `MediaXPlugin` 进行修改，以适配相关的插件需求。

## Form一个仓库

```shell
# 地址如下
https://github.com/ArtisanCloud/MediaXPlugin
```

## 修改步骤

### 1. 修改 `go.mod`

将 `module` 从原始仓库地址更改为新的仓库地址：

```diff
-module github.com/ArtisanCloud/MediaXPlugin
+module github.com/ArtisanCloud/MediaXPlugin-Wechat
```

### 2. 修改 `plugins.yaml`

原始 `plugins.yaml` 文件内容：

```yaml
name: PluginMediaX
version: 1.0.0
type: open
plugins:
  - name: PluginMediaXA
    sourcePath: ./src/pluginA.go
    buildPath: ./plugins/pluginA.so
    config: {}

  - name: PluginMediaXB
    sourcePath: ./src/pluginB.go
    buildPath: ./plugins/pluginB.so
    config: {}
```

如果你有两个插件，在一个插件bundle里，修改后：

```yaml
name: MediaXPluginWechat
version: 1.0.0
type: open
plugins:
  - name: WechatChannel
    sourcePath: ./src/channel.go
    buildPath: ./plugins/channel.so
    config: {}

  - name: WechatOfficialAccount
    sourcePath: ./src/officialAccount.go
    buildPath: ./plugins/officialAccount.so
    config: {}
```



## 3. 生成一个插件 MediaXPluginWechatOfficialAccount
```shell
cp src/pluginA.go src/officialAccount.go

```

修改officialAccount.go

* 全文件替换MediaXPlugin -> MediaXPluginWechatOfficialAccount

* 修改PluginName
```go

func NewMediaXPluginWechatOfficialAccount() MediaXPluginWechatOfficialAccount {
	return MediaXPluginWechatOfficialAccount{
		// PluginName: "PluginMediaX"
		PluginName: string(contract.WechatOfficialAccount),
	}
}

```
* 修改暴露插件的变量对象
```go
// var PluginMediaXA MediaXPluginRedBook = NewMediaXPluginRedBook()
var WechatOfficialAccount MediaXPluginWechatOfficialAccount = NewMediaXPluginWechatOfficialAccount()

```


### 4. 修改 `Makefile`

帮助编译插件组件，使用makefile工具

原始 `Makefile`:

```makefile
plugin.all.build: plugin.pluginA.build plugin.pluginB.build

.PHONY: plugin.pluginA.build
plugin.pluginA.build:
	@echo "正在构建pluginA插件..."
	go build -o plugins/pluginA.so -buildmode=plugin src/pluginA.go
	@echo "pluginA插件构建完成"

.PHONY: plugin.pluginB.build
plugin.pluginB.build:
	@echo "正在构建pluginB插件..."
	go build -o plugins/pluginB.so -buildmode=plugin src/pluginB.go
	@echo "pluginB插件构建完成"
```

修改后：

```makefile
plugin.all.build: plugin.officialAccount.build plugin.channel.build

.PHONY: plugin.officialAccount.build
plugin.officialAccount.build:
	@echo "正在构建officialAccount插件..."
	go build -o plugins/officialAccount.so -buildmode=plugin src/officialAccount.go
	@echo "officialAccount插件构建完成"

.PHONY: plugin.channel.build
plugin.channel.build:
	@echo "正在构建channel插件..."
	go build -o plugins/channel.so -buildmode=plugin src/channel.go
	@echo "channel插件构建完成"
```

## 运行构建

执行以下命令构建所有插件：

```sh
make plugin.all.build
```

