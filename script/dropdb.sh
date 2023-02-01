#!/usr/bin/env bash

VIEWS=`psql ${DATABASE_URI} -t --command "SELECT string_agg(table_name, ',') FROM information_schema.tables WHERE table_schema='public' AND table_type='VIEW'"`
BASETBLS=`psql ${DATABASE_URI} -t --command "SELECT string_agg(table_name, ',') FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE'"`
TYPES=`psql ${DATABASE_URI} -t --command "SELECT string_agg(distinct pg_type.typname, ',') FROM pg_type JOIN pg_enum ON pg_enum.enumtypid = pg_type.oid"`

echo Dropping views:${VIEWS}
psql ${DATABASE_URI} --command "DROP VIEW IF EXISTS ${VIEWS} CASCADE"
echo Dropping tables:${BASETBLS}
psql ${DATABASE_URI} --command "DROP TABLE IF EXISTS ${BASETBLS} CASCADE"
echo Dropping user defined types:${TYPES}
psql ${DATABASE_URI} --command "DROP TYPE IF EXISTS ${TYPES} CASCADE"