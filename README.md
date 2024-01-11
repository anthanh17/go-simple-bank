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

# Unit test

```
make test
```

# How to run

```
# Setup database
make postgres  # 1. create container docker postgres
make createdb  # 2. create a new database simple_bank in postgres-docker
```

3. Then connect to it using `TablePlus application`

```
make migrateup # 4. create structure simple_bank db
make server # 5. Run server
```
