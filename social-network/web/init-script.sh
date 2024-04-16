#!/bin/sh
# Replace API_HOST placeholder with actual environment variable
sed -i 's|http://localhost:39006|'"$API_HOST"'|g' /usr/share/nginx/html/static/js/*.js
