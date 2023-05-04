#!/bin/sh

set -e
COMMAND=$@

mkdir -p ~/.local/c2chopper/plugins
# curl ZIPPED_PLUGINS_URL -o ~/.local/c2chopper/plugins/plugins.zip
# tar -xzvf ~/.local/c2chopper/plugins/plugins.zip --directory ~/.local/c2chopper/plugins/plugins.zip
# unzip ~/.local/c2chopper/plugins/plugins.zip -d ~/.local/c2chopper/plugins/

exec $COMMAND
