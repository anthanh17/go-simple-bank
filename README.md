# go-simple-bank
## Function:
1. Create and manage account: Owner, balance, currency
2. Record all balance changes: Create an account entry for each change
3. Money transfer transaction: Perform money transfer between 2 accounts consistently within a transaction
## Database design:
- Design database schema: Design a SQL DB using dbdiagram.io
- Generate SQL code to create the schema in target database engine: Postgress, MySQL
## Setup:
- wsl2
- docker
- go
- vscode
- make
- sqlc: genarate golang codes from SQL queries
- golang-migrate: migrate database `brew install golang-migrate`
## Use Docker:
```
docker run
--name <container_name>
-e <environment_variable>
-p <host_ports:container_ports>
-d
<image>:<tag>
```
```
docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=abc123 -d postgres:12-alpine
```
Run command in container
```
docker exec -it
<container_name_or_id>
<command> [args]
```
```
docker exec -it postgres12 psql -U root
# comand run without password because container is trust
# or
docker exec -it postgres12 /bin/sh
```
Internal postgres12
```
select now();
\q
```
View container logs
```
docker logs
<container_name_or_id>
```
## How to write & run database migration in Golang
Usage:
```
migrate -help

- create: Can use to create new migration files
- goto: Will migrate the scheme ti a specific version
- up/down: Apply all or N up/down migrations
```
Start with
```
migrate create -ext sql -dir db/migration -seq init_schema

```
## SQLC
- Have many choose `database/sql | gorm | sqlx | sqlc`
    - Prioritize: sqlc > sqlx (fast ? failure won't occur until runtime) > gorm (slow) > database/sql (easy make mistakes)
- Very fast & easy to use
- Automatic code genaretion `gen golang using standard lib database/sql => fast`
- Catch SQL query errors before genarating code
- Full support Postgres. MySQL is experimental
### Install sqlc
```
brew install sqlc # macos
sudo snap install sqlc # ubuntu
```
Usage:
```
sqlc help

- compile: Statically check SQL for syntax and type errors
- generate: Generate source code from SQL
- init: Create an empty sqlc.yaml settings file
```
# ----------------------------------------------
# How to run
```
# Setup database
make postgres # create container docker postgres
make createdb # create db simple_bank in docker
make migrateup # create structure simple_bank db

# Setup sqlc
sqlc init
```
