#!/bin/bash
set -e
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE frontdemo;
    CREATE USER doseok WITH ENCRYPTED PASSWORD 'kim7795004';
    GRANT ALL PRIVILEGES ON DATABASE frontdemo TO doseok ;
EOSQL