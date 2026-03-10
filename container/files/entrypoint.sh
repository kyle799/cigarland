#!/bin/sh
set -e

echo "Waiting for postgres at ${DB_HOST}:${DB_PORT:-5432}..."
until nc -z "${DB_HOST}" "${DB_PORT:-5432}"; do
  echo "Postgres not ready, retrying..."
  sleep 1
done
echo "Postgres ready."

exec /cigarland/cigarland_api -start-server -create-db
