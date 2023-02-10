import React from 'react';

class PostWillRenderEmbed extends React.Component {
    static plugin = null;

    render() {
        let title = '';
        const iframeWidth = PostWillRenderEmbed.plugin.props.iframeWidth;
        const iframeHeight = PostWillRenderEmbed.plugin.props.iframeHeight;

        try {
            title = this.props.embed.data.title;
        } catch {
        }

        let url = '';
        if (!this.props.embed.url.includes('player.vimeo.com')) {
            url = this.props.embed.url.replace('vimeo.com', 'player.vimeo.com/video');
        } else {
            url = this.props.embed.url;
        }

        return (
            <div>
                <h5>{title}</h5>
                <iframe src={url} width={iframeWidth} height={iframeHeight}>
                </iframe>
            </div>
        );
    }
}

class VimeoPlugin {
    initialize(registry, store) {
        const plugin = store.getState().plugins.plugins.vimeo;
        PostWillRenderEmbed.plugin = plugin;
        registry.registerPostWillRenderEmbedComponent(
            (embed) => {
                if (embed.type == 'opengraph' && embed.url.includes('vimeo.com')) {
                    return true;
                }
                return false;
            },
            PostWillRenderEmbed,
            false,
        );
    }

    uninitialize() {
        // No clean up required.
    }
}

window.registerPlugin('vimeo', new VimeoPlugin());