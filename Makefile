all:
	make prepare
	make linux
	make macos
	make windows
	make webapp
	make pack

prepare:
	rm -f mattermost-vimeo-plugin.tar.gz
	rm -rf mattermost-vimeo-plugin
	mkdir -p mattermost-vimeo-plugin
	mkdir -p mattermost-vimeo-plugin/client
	mkdir -p mattermost-vimeo-plugin/server

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mattermost-vimeo-plugin/server/plugin-linux-amd64 server/plugin.go

macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o mattermost-vimeo-plugin/server/plugin-darwin-amd64 server/plugin.go

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o mattermost-vimeo-plugin/server/plugin-windows-amd64 server/plugin.go

webapp:
	mkdir -p dist
	npm install
	./node_modules/.bin/webpack --mode=production

pack:
	cp -r dist/main.js mattermost-vimeo-plugin/client
	cp plugin.json mattermost-vimeo-plugin/
	tar -czvf mattermost-vimeo-plugin.tar.gz mattermost-vimeo-plugin
