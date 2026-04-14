#!/bin/bash
set -euo pipefail

SOURCE_URL=https://github.com/Angus-Warman/hello-world-complete/releases/latest/download/hello-world-complete_linux_amd64
TEMP=/home/azureuser/hello-world-complete/hello-world-complete_linux_amd64-TEMP
FINAL=/home/azureuser/hello-world-complete/hello-world-complete_linux_amd64
SCREEN_NAME=hello-world-complete

echo "Downloading latest release..."
curl -L $SOURCE_URL -o $TEMP
echo "Closing previous version..."
screen -S $SCREEN_NAME -X quit || true
mv $TEMP $FINAL
chmod +x $FINAL
echo "Starting server..."
screen -dmS $SCREEN_NAME $FINAL
echo "Done"