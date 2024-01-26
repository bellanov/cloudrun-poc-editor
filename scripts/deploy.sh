#!/bin/bash
#
# Execute Docker Container.
GOOGLE_APPLICATION_CREDENTIALS="/home/bella/.config/gcloud/application_default_credentials.json"

sudo docker run -e PORT --env PORT=3000 \
    -e EDITOR_UPSTREAM_RENDER_URL --env EDITOR_UPSTREAM_RENDER_URL=3001 \
    -e GOOGLE_APPLICATION_CREDENTIALS --env GOOGLE_APPLICATION_CREDENTIALS=/tmp/keys/credentials.json \
    -v $GOOGLE_APPLICATION_CREDENTIALS:/tmp/keys/credentials.json:ro \
    -p 3000:3000 editor:local
