#!/bin/sh

### DATABASE
# pull postgres image
docker pull postgres:18
# run the container with a volume
# todo: get pwd from debug.env
docker run --name postgres-db -v postgres-volume:/var/lib/postresql/18/docker --env-file debug.env postgres:18

# docker build
# docker run db
# docker run orderservice
docker build -t orderservice-img .
docer run --name orderservice orderservice-img

# create network for them containers to communicate