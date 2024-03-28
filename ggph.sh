#!/bin/sh

USERNAME="imanaspaul"
REPO="ggph"
BINARY_NAME="ggph"
VERSION="v0.0.1"

# Download the binary from GitHub release
wget -O $BINARY_NAME https://github.com/$USERNAME/$REPO/releases/download/$VERSION/$BINARY_NAME
                    #  https://github.com/imanaspaul/ggph/releases/download/v0.0.1/ggph


chmod +x $BINARY_NAME

./$BINARY_NAME