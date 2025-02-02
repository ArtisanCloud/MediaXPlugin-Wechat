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