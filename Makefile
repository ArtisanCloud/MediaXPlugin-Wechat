# 单独构建插件文件，传入模块名
.PHONY: plugin.build
plugin.build:
	@echo "正在构建插件..."
	go build -o plugin/plugin.so -buildmode=plugin src/plugin.go
	@echo "插件构建完成"