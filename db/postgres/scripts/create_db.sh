#!/bin/bash

# -- CREATE DB -- #
# Note:
# It's OK to run this on an existing db, it will fail gracefully
#

# VARIABLES NEED TO PASS IN -- #
# DATABASE_HOST - database host address
# DATABASE_USERNAME -- master datebase user
# DATABASE_PASSWORD -- master database password
# DATABASE_NAME - creating database name
# DATABASE_USERNAME_RO - creating read only user for the new database
# DATABASE_USERNAME_RW - creating write only user for the new database


# use the master db credentials
export PGPASSWORD=$DATABASE_PASSWORD

set -x

# create the database
createdb -h $DATABASE_HOST -p 5432 -U $DATABASE_USERNAME -w $DATABASE_NAME

# create ro and rw users
createuser -h $DATABASE_HOST -p 5432 -U $DATABASE_USERNAME --no-createdb --no-superuser --no-replication -w $DATABASE_USERNAME_RO
createuser -h $DATABASE_HOST -p 5432 -U $DATABASE_USERNAME --createdb --no-superuser --no-replication -w $DATABASE_USERNAME_RW

# stop logging to mask passwords
set +x

# update passwords
echo "ALTER USER \"$DATABASE_USERNAME_RO\" WITH PASSWORD '$DATABASE_PASSWORD_RO';" \
    | psql -qtAX --host=$DATABASE_HOST --dbname=postgres --username=$DATABASE_USERNAME
echo "ALTER USER \"$DATABASE_USERNAME_RW\" WITH PASSWORD '$DATABASE_PASSWORD_RW';" \
    | psql -qtAX --host=$DATABASE_HOST --dbname=postgres --username=$DATABASE_USERNAME

echo "passwords updated"

# fix extra permissions
echo "ALTER DEFAULT PRIVILEGES FOR ROLE \"$DATABASE_USERNAME_RW\" GRANT SELECT ON TABLES TO \"$DATABASE_USERNAME_RO\";" \
    | PGPASSWORD=$DATABASE_PASSWORD_RW psql -qtAX --host=$DATABASE_HOST --dbname=$DATABASE_NAME --username=$DATABASE_USERNAME_RW

echo "permissions updated"
