#!/bin/sh

command -v sqlx > /dev/null || cargo install sqlx-cli --no-default-features --features postgres,native-tls

export PROJECT_ROOT=$(dirname $(dirname $(realpath $0)))
export DATABASE_URL=${MANEKANI_DATABASE_URL:-"postgres://manekani:secret@postgres-manekani/manekani-test"}

sqlx database create
sqlx migrate run --source "${PROJECT_ROOT}"/db