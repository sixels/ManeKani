#!/bin/sh

docker compose -f docker-compose.yml \
  -f infra/auth/compose.yml \
  -f infra/db/compose.yml \
  -f docker-compose.override.yml \
  $@
