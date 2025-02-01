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