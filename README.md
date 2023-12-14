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

# How to run

```
# Setup database
make postgres # create container docker postgres
make createdb # create db simple_bank in docker
make migrateup # create structure simple_bank db

# Setup sqlc
sqlc init
# config sqlc.yaml
# define sql in folder db/query
make sqlc # gen code golang from sql

# unit test
make test
```
