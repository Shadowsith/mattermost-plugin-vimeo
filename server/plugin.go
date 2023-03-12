package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mattermost/mattermost-server/v6/plugin"
)

// VimeoPlugin implements the interface expected by the Mattermost server to communicate
// between the server and plugin processes.
type VimeoPlugin struct {
	plugin.MattermostPlugin
}

type PluginSettings struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

const (
	routeSettings = "/settings"
)

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *VimeoPlugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	settings := p.getSettings(w)
	res := p.handleSettingsResult(settings)
	fmt.Fprint(w, res)
}

func (p *VimeoPlugin) getSettings(w http.ResponseWriter) PluginSettings {
	pluginSettings, ok := p.API.GetConfig().PluginSettings.Plugins["vimeo"]
	settings := PluginSettings{
		Height: 340,
		Width:  600,
	}
	if ok {
		for k, v := range pluginSettings {
			switch k {
			case "height":
				settings.Height = p.getIntVal(v)
				break

			case "width":
				settings.Width = p.getIntVal(v)
				break
			}
		}
	}
	return settings
}

func (p *VimeoPlugin) getIntVal(v interface{}) int {
	val, ok := strconv.Atoi(fmt.Sprintf("%v", v))
	if ok != nil {
		val = 20
	}
	return val
}

func (p *VimeoPlugin) handleSettingsResult(settings PluginSettings) string {
	json, err := json.Marshal(&settings)
	if err != nil {
		return "{}"
	} else {
		return string(json)
	}
}

// This example demonstrates a plugin that handles HTTP requests which respond by greeting the
// world.
func main() {
	plugin.ClientMain(&VimeoPlugin{})
}
