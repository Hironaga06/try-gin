#!/bin/bash

set -e

SCHEMA_PATH=$1
TABLE_NAME=$2

echo "===== create migration files ====="
echo "$(date +%Y%m%d%H%M%S)_$TABLE_NAME.up.sql"
touch $SCHEMA_PATH/$(date +%Y%m%d%H%M%S)_$TABLE_NAME.up.sql
echo "$(date +%Y%m%d%H%M%S)_$TABLE_NAME.down.sql"
touch $SCHEMA_PATH/$(date +%Y%m%d%H%M%S)_$TABLE_NAME.down.sql

echo "finish"