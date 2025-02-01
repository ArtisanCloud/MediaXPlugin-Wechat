plugin.all.build: plugin.officialAccount.build plugin.moments.build

.PHONY: plugin.officialAccount.build
plugin.officialAccount.build:
	@echo "正在构建officialAccount插件..."
	go build -o plugins/officialAccount.so -buildmode=plugin src/officialAccount.go
	@echo "officialAccount插件构建完成"

.PHONY: plugin.moments.build
plugin.moments.build:
	@echo "正在构建moments插件..."
	go build -o plugins/moments.so -buildmode=plugin src/moments.go
	@echo "moments插件构建完成"