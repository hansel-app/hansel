#!/bin/bash
set -e

# Replace placeholders in env.js with values from the environment
envsubst < /build/env.template.js > /build/env.js
rm -f /build/env.template.js

# Then exec the container's main process (what's set as CMD in the Dockerfile)
exec "$@"
