#!bin/bash

set -e

echo "- waiting for postgres"
until psql ${DATABASE_URL} -c '\q'; do
    sleep 1
done

echo "- exec migrate up"
go run ./script/migration_databases.go -exec up

echo "- finish!"