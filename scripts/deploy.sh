#!/bin/bash
set -euo pipefail

SOURCE_URL=https://github.com/Angus-Warman/hello-world-complete/releases/latest/download/hello-world-complete_linux_amd64
DIR=/home/azureuser/hello-world-complete/
EXEC_NAME=hello-world-complete_linux_amd64
SCREEN_NAME=hello-world-complete
TEMP="${DIR}${EXEC_NAME}-TEMP"
FINAL="${DIR}${EXEC_NAME}"

rm -f $TEMP
echo "Downloading latest release..."
curl --no-progress-meter -L $SOURCE_URL -o $TEMP
echo "Closing previous version..."
pkill -f "$EXEC_NAME" || true
mv -f $TEMP $FINAL
chmod +x $FINAL
echo "Starting server..."
cd $DIR # Go to folder to get correct .env
screen -dmS $SCREEN_NAME $FINAL
echo "Done"
