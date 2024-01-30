How to develop

# Setup

### 1. Setup docke database postgress

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

### 2. How to write & run database migration in Golang

Install:

```
brew install golang-migrate
migrate -version
```

Usage:

```
migrate -help

- create: Can use to create new migration files
- goto: Will migrate the scheme ti a specific version
- up/down: Apply all or N up/down migrations
```

```
# step1: create new migration files
migrate create -ext sql -dir db/migration -seq init_schema
# step 2: apply all up migration
migrate -path db/migration -database "postgresql://root:abc123@localhost:5432/simple_bank?sslmode=disable" -verbose up

```

Then we can use migaration file to up/down schema sql

- First step: copy all sql file to init_schema_up.sql and write drop table code to init_schema_down.sql (revert the changes made by the init_schema_up.sql)

### 3. Install sqlc

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

## Why choose use to SQLC

- Have many choose `database/sql | gorm | sqlx | sqlc`
  - Prioritize: sqlc > sqlx (fast ? failure won't occur until runtime) > gorm (slow) > database/sql (easy make mistakes)
- Very fast & easy to use
- Automatic code genaretion `gen golang using standard lib database/sql => fast`
- Catch SQL query errors before genarating code
- Full support Postgres. MySQL is experimental

## How to use sqlc

```
# Setup sqlc
sqlc init
# config sqlc.yaml
# define sql in folder db/query
make sqlc # gen code golang from sql
```

1. sqlc init # file config sqlc -> create sqlc.yaml
2. create `schema.sql` in this case is \*.sql file in "./db/migration/"

```
# Example
CREATE TABLE authors (
  id   BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name text    NOT NULL,
  bio  text
);
```

3. create `query.sql`

```
# Example
-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = ? LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :execresult
INSERT INTO authors (
  name, bio
) VALUES (
  ?, ?
);

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = ?;
```

4. Final: run command gen code golang

```
sqlc generate
```

## Unit testing Go database CRUD

- In Golang, we have a convention to put test file in the same folder with the code
- And the name of the test file should end with the `test suffix`
- In the same folder: account.sql.go and account_test.go

=============================================================

# A clean way to implement database transaction in Golang

## What is a db transaction?

- A single unit of work
- Often made up of multiple db operations
- Ex: In our simple bank, we want to transfer `10 USD` from `account 1` to `account 2`
  - This transaction comprises 5 operations:
    - `1`.First, Create a transfer record with amount = 10
    - `2`.Secound, Create an account entry record for `account 1` with amount = `-10`, since money is moving out of this account.
    - `3`.Create another account entry record for `account 2` with amount = `+10`, because money is moving in to this account.
    - `4`.Then we update the balance of `account 1` by `subtracting 10` from it
    - `5`.And finally we update the balance of `account 2` by `adding 10` to it
    - `=> This is transaction that we're going to implement.`

## Why do we need db transaction?

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

## How to run SQL database transaction?

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

# How to avoid Db deadlock?

One of the hardest thing when working with database transaction is `locking` and `handling deadlock`

=> The best way to deal with dealock is to avoid it.

=> By that i mean we should fine-tune our queries in the transaction so that `deadlock won't have a chance to occur` or `at least minimize its chance of occurence.`

=> avoid them by making sure that `out application always acquire locks in a consistent order`

# Mock DB for testing HTTP API in Go

## Why mock database?

1. It helps us to write independent tests more easily because each test will use its own separate mock db to store data, so there will be no conflicts between them.
   > If you use a real db, all tests will read and write data to the same place, so it would be harder to avoid conflicts, especially in a big project with a large code base.
2. Our tests will run much faster since they don't have to spend time talking to the database and waiting for the queries to run.
   > All actions will be performed in memory, and within the same process.
3. Very important reason for mocking database is: `It allows us to write tests that achieve 100% coverage`
   > With a mock db, we can easily setup and test some edge cases, such as an unexpected error, or a connection lost, which would be impossible to achieve if we use a real db.

## Is it good enough to test our API with a mock DB?

`Yes absolutely!`

Because our code that talks to the real DB is already tested carefully in the previous lectures.

-> So all we need to do is: make sure that the mock db implement the same interface as the real db -> The everything will be working just fine when being put together.

## How to mock?

There are 2 ways to mock db:

1. `USE FAKE DB` Implement a fake version of DB: store data un memory
2. `USE DB STUBS: GOMOCK` Generate and build stubs that returns hard-coded values

-> Use th the secound way

## How to use

```
# 0. Install
go get github.com/golang/mock/mockgen@v1.6.0
# 1. Gen code
mockgen -package mockdb -destination db/mock/store.go github.com/anthanh17/simplebank/db/sqlc Store
```

# How to securely store passwords? Hash password in Go with Bcrypt!

`We should never ever store naked password!`

-> Idea: hash it first & only store that hash value

1. The password will be hashed using `brypt hashing function` to produce a hash value. Besides `the input password` bcrypt requires a `cost` parameter and `salt` parameter.

- Cost: this value which will decide the number of key expansion rounds or iterations of the algorithm
- Salt: this value is random, to be used in those iterations => which will help protect against the rainbow table attack

-> Then the string will store in the database

- When users login, how can we verify that the passoword that they entered is correct or not?
  > 1. First we have to find the hashed password stored in the DB `by username`
  > 2. Then we use `cost and salt` of that hashed passoword as the arguments to hash `the password users just entered` with bcrypt => the string
  > 3. Compare the 2 hash values => verify

# JWT

- The string is base64-encoded

- 3 main parts and separation by "."

1. Header:

```
{
  "alg:: "HS256",
  "typ": "JWT"
}
```

- Decode this part: JSON object contains type and algorithm used to sign the token.

2. Payload:

```
{
  ...
}
```

- Where we store information
  > Keep in mind: That all data stored in JWT is only base64-encoded, not encrypted.
  > => Don't need the secret/private key of the server in order to decode its content

3. Verify signature

- The idea: Only the server has the secret/private key to sign the token
  > So if hacker attempts to create a fake token without the correct key -> Server will be easily detected in the verification process

## JWT signing algorithms

### 2 main categories:

1. Symmetric-key algorithm: Thuật toán khóa đối xứng

   > Trong đó cùng 1 secret key được sử dụng để ký và xác minh token

   > Và vì chỉ có 1 key nên nó cần được giữ bí mật => Thuật toán này chỉ phù hợp để sử dụng cục bộ. (Nghĩa là dùng cho các service nội bộ , nơi secret key có thể được chia sẻ)

   > HS256 = HMAC + SHA 256, HS384, HS512 2.

   > Thuật toán khóa đối xứng hiệu quả và phù hợp với hầu hết các ứng dụng. Tuy nhiên, không thể sử dụng nó trong trường hợp có service bên ngoài muốn verify token

2. Asymmetric digital signature algorithm: Thuật toán khóa bất đối xứng

   > Trong thuật toán này, `có 1 cặp key` thay vì chỉ 1 key secret duy nhất

   > `private key` sử dụng để ký token, `public key` dùng để verify token

   > => Từ đó ta có thể dễ dàng share `public key` với bất kì service bên thứ 3 bên ngoài mà không sợ rò rỉ private key

   > RS group, PS group, ES group: RS256, RS384, RS512 || PS256, PS384, PS512 || ES, ES. ES.

## What's the problem of JWT

1. Weak algorithms:
   > JWT cung cấp quá nhiều các thuật toán để dev lựa chọn, bao gồm cả các thuật toán đã được biết là dễ bị tấn công (RSA PKCSv1.5 dễ bị tấn công) -> với dev ít exp sẽ không biến nên sử dụng thuật toán nào là tốt nhất
2. Trivial Forgery:
   > JWT làm cho việc giả mạo token trở nên tầm thường đến mức nếu bạn không cẩn thận trong quá trình triển khai or chọn thư viện triển khai kém cho prj của mình -> Hệ thống của bán sẽ dễ dàng trở thành mục tiêu dễ bị tấn công
   > JWT nó bao gồm thuật toán ký trong header token -> hacker chỉ cần set trường "alg" trong header thành "none" để vượt qua quá trình verify chữ kí -> 1 số thư viện đã fixed vấn đề này -> đây là điểu nên cẩn thận
   > 1 vấn đề nữa là hacker cố tình set thuật toán trong header thành đối xứng ví dụ: HS256 trong khi biết rằng server thực sự sử dụng thuật toán bất đối xứng

=> JWT không phải là 1 tiêu chuẩn được thiết kê tốt.
=> Sử dụng `PASETO - Platform-Agnostic Security Tokens`

# PASETO - Platform-Agnostic Security Tokens

1. Stronger algorithms:
   > dev không phải chọn thuật toán nữa. Họ chỉ cần chọn phiên bản PASETO mà muốn dùng. Mỗi phiên bản được triển khai với 1 bộ mật mã mạn
2. Non-trivial Forgery

   > Header thuật toán không còn tồn tại nữa -> hacker không thế đặt nó thành "none" or buộc server sử dụng thuật toán mà nó đã chọn trong header.

   > Everything is authenticated

   > Encrypted payload for local use <symmetric key>

3. An toàn và đơn giản hơn JWT

# Build a minimal Golang Docker image with a multistage Dockerfile

## How to deloy an application

1. Build image
2. ship it -> Deloy the app to AWS

`Bước 1:` Dockerize the application

- Create Dockerfile

```
# 1. Xác định image cở sở để build app
FROM golang:1.16-alpine3.13

# 2. Sử dụng lệnh WORKDIR để khai báo thư mục làm việ hiện tại bên trong image
WORKDIR /app

# 3. Sau đó copy all các file cần thiết vào folder này
COPY . .

# 4. Chúng ta sẽ build ứng dụng của mình thành binary executable file.
RUN go build -o main main.go

# 5. Cách tốt nhất cũng là sử dụng EXPOSE để thông báo cho Docker rằng container sẽ lắng nghe trên cổng mạng được chỉ định trong runtime.
EXPOSE 9000

# 6. Xác định lệnh mặc định sẽ chạy khi container run
CMD ["/app/main"]
```

```
# Build dockerfile
docker build -t simplebank:latest .

docker images
# Run container
docker run --name simplebank -p 9000:9000 -e GIN_MODE=release simplebank:latest
```

> Ở đây nếu build ra sẽ thấy image khá nặng tầm 456mb

> Khắc phục bằng cách sửa Dockerfile: write a lulti-stage Dockerfile để giảm dung lượng image xuôgns còn tầm 17.6mb
