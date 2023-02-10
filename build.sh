./node_modules/.bin/webpack --mode=production
rm -f mattermost-vimeo-plugin.tar.gz
rm -rf mattermost-vimeo-plugin
mkdir -p mattermost-vimeo-plugin
cp -r dist/main.js mattermost-vimeo-plugin/
cp plugin.json mattermost-vimeo-plugin/
tar -czvf mattermost-vimeo-plugin.tar.gz mattermost-vimeo-plugin