#!/usr/bin/env bash

set -e

# source ./hack/lib/common.sh

# header_text "Running markdown link checker against changed markdown files"
docker run --rm -i -v "$(pwd):/src" -v "$(pwd)/website/public:/target" klakegg/hugo:0.62.2-ext-ubuntu -s website
docker run -v "$(pwd)/website/public:/target" mtlynch/htmlproofer /target --empty-alt-ignore --http-status-ignore 429
