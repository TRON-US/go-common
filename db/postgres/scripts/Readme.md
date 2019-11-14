# postgres database generic scripts

## create_db

#### Goal:

1. create db
2. create rw role
3. create ro role
4. alter default privileges

#### Notice:

* It's OK to run this on an existing db, it will fail gracefully

* Variables need to pass in
```
VARIABLES NEED TO PASS IN
DATABASE_HOST - database host address
DATABASE_USERNAME -- master datebase user
DATABASE_PASSWORD -- master database password
DATABASE_NAME - creating database name
DATABASE_USERNAME_RO - creating read only user for the new database
DATABASE_USERNAME_RW - creating write only user for the new database
```

## drop_db

#### Goal:

1. disconnect all connections
2. drop privileges
3. drop roles
4. drop db

#### Notice:

* Variables need to pass in
```
DATABASE_HOST - database host address
DATABASE_USERNAME -- master datebase user
DATABASE_PASSWORD -- master database password
DATABASE_NAME - droping database name
DATABASE_USERNAME_RO - read only user for the droping database
DATABASE_USERNAME_RW - write only user for the droping database
```