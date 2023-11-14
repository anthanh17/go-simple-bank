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
```
View container logs
```
docker logs
<container_name_or_id>
```
