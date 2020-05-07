#!/usr/bin/env bash

set -e

source ./hack/lib/common.sh

header_text "Building html and checking links"

klakegg/hugo:0.69.2-ext-ubuntu

pushd website
npm install postcss-cli autoprefixer
hugo
popd
