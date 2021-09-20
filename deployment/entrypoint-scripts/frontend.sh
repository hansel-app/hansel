#!/bin/bash
set -e

# Build app
npm run build

# Then exec the container's main process (what's set as CMD in the Dockerfile)
exec "$@"
