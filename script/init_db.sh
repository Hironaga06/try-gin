#!bin/bash

set -e

echo "- waiting for postgres"
until psql 'postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable' -c '\q'; do
    sleep 1
done

echo "- exec migrate up"
go run ./script/migration_databases.go -exec up

echo "- finish!"

tail -f /dev/null