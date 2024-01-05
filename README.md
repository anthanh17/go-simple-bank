# go-simple-bank

## Function:

1. Create and manage account: Owner, balance, currency
2. Record all balance changes: Create an account entry for each change
3. Money transfer transaction: Perform money transfer between 2 accounts consistently within a transaction

## Database design:

- Design database schema: Design a SQL DB using `dbdiagram.io`

-> Generate SQL code to create the schema in target database engine: Postgress, MySQL

## Setup:

- wsl2 (option with windows)
- go
- vscode
- docker
- make
- sqlc: genarate golang codes from SQL queries
- golang-migrate: migrate database `brew install golang-migrate`

# How to run

```
# Setup database
make postgres  # 1. create container docker postgres
make createdb  # 2. create db simple_bank in docker
make migrateup # 3. create structure simple_bank db

# Setup sqlc
sqlc init
# config sqlc.yaml
# define sql in folder db/query
make sqlc # gen code golang from sql

# unit test
make test
```
