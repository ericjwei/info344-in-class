#!/usr/bin/env bash
docker rm -f zipsvr

docker run -d \
-p 443:443 \
--name zipsvr \
-v /mnt/d/UW/Info344/go/src/github.com/ericjwei/info344-a17/info344-in-class/zipsvr/tls:
/tls:ro \
-e TLSCERT=/tls/fullchain.pem \ 
-e TLSKEY=/tls/privkey.pem \
ericjwei/zipsvr
