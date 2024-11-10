#!/usr/bin/env bash

set -eux

if [ -z $(docker network ls -f name=$POD_NETWORK --format "{{ .ID }}") ]; then
  docker network create $POD_NETWORK &> /dev/null
fi

if [ ! -z $(docker ps -a -f name=$CONTAINER_POD --format "{{ .ID }}") ]; then
  docker rm -f $POD_CONTAINER &> /dev/null
fi

POD_IMAGE=postgres:16.0
POD_CONTAINER=pg-debug-pod
POD_VOLUME=pg-debug-pod-data
POD_NETWORK=debug-pod-net
POD_PORT=5432

PGDATA=/var/lib/postgresql/data/pgdata
PGUSER=pgadmin
PGPASS=pgadmin
PGDBNAME=pgadmindb

NUM_CPUS=1
MEM_LIMIT=2000000000
docker run --rm -d -ti \
  --cpus $NUM_CPUS \
  --memory $MEM_LIMIT \
  --name $POD_CONTAINER \
  --network $POD_NETWORK \
  -e PGDATA=$PGDATA \
  -e POSTGRES_DB=$PGDBNAME \
  -e POSTGRES_USER=$PGUSER \
  -e POSTGRES_PASSWORD=$PGPASS \
  -v $POD_VOLUME:$PGDATA \
  -p $POD_PORT:5432 \
  $POD_IMAGE
