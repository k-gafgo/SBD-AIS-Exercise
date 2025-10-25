#!/bin/sh

### DATABASE
# pull postgres image
docker pull postgres:18
# run the container with a volume
# i exposed the db to port 5433 on my host because i already have postgresql installed which is listening on port 5432
docker run --name postgres-db -v postgres-volume:/var/lib/postgresql/data -p 5433:5432 --env-file debug.env postgres:18

# prepoluate database
export POSTGRES_DB=order
export POSTGRES_USER=docker
export POSTGRES_PASSWORD=docker
export POSTGRES_TCP_PORT=5433
export DB_HOST=127.0.0.1

go run main.go

### NETWORK
# create network for the containers to communicate
docker network create order-net
docker network connect order-net postgres-db

### ORDER SERVICE
docker build -t orderservice-img .
docker run --name orderservice --net order-net -p 3000:3000 --env-file debug.env orderservice-img