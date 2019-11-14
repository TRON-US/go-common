#!/bin/bash

# -- VARIABLES NEED TO PASS IN -- #
# DATABASE_HOST - database host address
# DATABASE_USERNAME -- master datebase user
# DATABASE_PASSWORD -- master database password
# DATABASE_NAME - droping database name
# DATABASE_USERNAME_RO - read only user for the droping database
# DATABASE_USERNAME_RW - write only user for the droping database

# use the master db credentials
export PGPASSWORD=$DATABASE_PASSWORD

# disconnect all connections
echo "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname=\"$DATABASE_NAME\";" \
    | psql -qtAX --host=$DATABASE_HOST --dbname=postgres --username=$DATABASE_USERNAME

# drop default privileges
echo "DROP OWNED BY \"$DATABASE_USERNAME_RO\";" \
    | PGPASSWORD=$DATABASE_PASSWORD_RO psql -qtAX --host=$DATABASE_HOST --dbname=postgres --username=$DATABASE_USERNAME_RO
echo "DROP OWNED BY \"$DATABASE_USERNAME_RW\";" \
    | PGPASSWORD=$DATABASE_PASSWORD_RW psql -qtAX --host=$DATABASE_HOST --dbname=postgres --username=$DATABASE_USERNAME_RW

set -x
dropdb -h $DATABASE_HOST -p 5432 -U $DATABASE_USERNAME -w $DATABASE_NAME
dropuser -h $DATABASE_HOST -p 5432 -U $DATABASE_USERNAME  $DATABASE_USERNAME_RO
dropuser -h $DATABASE_HOST -p 5432 -U $DATABASE_USERNAME $DATABASE_USERNAME_RW
