#!/usr/bin/env bash

set -eax

if [ -z $GEN_DIR]; then
  GEN_DIR=./gen
fi

if [ ! -d $GEN_DIR ]; then
  echo "Directory $GEN_DIR not found"
fi

GEN_HASH_DIR=$(find $GEN_DIR -type f | sha1sum | awk '{print $1}')
go generate ./...
REGEN_HASH_DIR=$(find $GEN_DIR -type f | sha1sum | awk '{print $1}')

if [ $GEN_HASH_DIR != $REGEN_HASH_DIR ]; then
  echo "Check gen hash failed"
fi
