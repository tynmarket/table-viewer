#!/bin/bash

echo "$AUTH_USER:$(openssl passwd $AUTH_PASSWORD)" > /etc/nginx/.htpasswd

nginx
./server
