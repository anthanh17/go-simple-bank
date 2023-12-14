How to develop

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

# Command exc db

```
docker exec -it postgres12 psql -U root
\l
\c simple_bank
\dt
select * from accounts;
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

# What is a db transaction?

- A single unit of work
- often made up of multiple db operations
- Ex: We want to transfer `10 USD` from `account 1` to `account 2`
  - This transaction comprises 5 operations:
    - `1`.First, Create a transfer record with amount = 10
    - `2`.Secound, Create an account entry for account 1 with amount = -10, since money is moving out of this account.
    - `3`.Create another account entry record for account 2 with amount = +10, because money is moving in to this account.
    - `4`.Then we update the balance of account 1 by subtracting 10 from it
    - `5`.And finally we update the balance of account 2 by adding 10 to it
    - `=> This is transaction that we're going to implement.`

# Why do we need db transaction?

There are 2 main reasons:

1. To provide a reliable and consistent unit of work, even in case of system failure
2. To provide isolation between programs that access the database concurrently

`In order to achieve these 2 goals, a database transaction must satisfy the ACID properties`

- `Atomicity (A)`:
  - Which means either all operations of the transaction complete successfully or the whole transaction fails, and `everything is rolled back` and the database is unchanged.
- `Consistency (C)`:
  - which means the database state should remains valid after the transaction is executed. `More precisely, all data written to the database must be valid` according to predefined rules, including constraints, cascades and triggers.
- `Isolation (I)`:
  - Concurrent transactions must not affect each other
- `Durability (D)`:
  - Which means that all data written by a successful transaction must stay in a persistent storage and cannot be lost, even in case of system failure

# How to run SQL database transaction?

```
BEGIN;
... # write a series of normal SQL queries
COMMIT;
```

If any query fails:

```
BEGIN;
... # write a series of normal SQL queries
ROLLBACK;
```
