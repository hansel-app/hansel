# When searching for a virtual server by name, if name matches more than one of the specified variants, e.g.
# both wildcard name and regular expression match, the first matching variant will be chosen, in the following
# order of precedence:
#
# 1. exact name
# 2. longest wildcard name starting with an asterisk, e.g. “*.example.org”
# 3. longest wildcard name ending with an asterisk, e.g. “mail.*”
# 4. first matching regular expression (in order of appearance in a configuration file)

server {
  listen 80;
  server_name hansel-app.com;

  # Redirect HTTP to HTTPS
  location / {
    return 301 https://$host$request_uri;
  }
}

server {
  listen 80;
  server_name www.hansel-app.com;

  # Redirect www to non-www
  return 301 http://hansel-app.com$request_uri;
}

server {
  listen 443 ssl http2;
  server_name hansel-app.com;
  root /build;
  index index.html;

  add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;

  ssl_certificate /etc/letsencrypt/live/hansel-app.com/fullchain.pem;
  ssl_certificate_key /etc/letsencrypt/live/hansel-app.com/privkey.pem;
  include /etc/letsencrypt/options-ssl-nginx.conf;
  ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;

  location / {
    # Pass on URL handling to vue-router
    try_files $uri $uri/ /index.html;
    # Disable caching of index.html so that any changes to the Vue application invalidates the cache
    # Note that the Vue application itself is still cached
    add_header Cache-Control "no-store, no-cache, must-revalidate";
  }
}

# Catch-all for unrecognised requests
server {
  listen 80 default_server;
  server_name _;
  return 444;
}
