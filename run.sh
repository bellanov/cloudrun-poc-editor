export EDITOR_UPSTREAM_RENDER_URL="http://renderer.local"
docker build -t editor .
docker run -p 8080:8080 -e EDITOR_UPSTREAM_RENDER_URL --rm --name editor editor