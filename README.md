
# Waizly Developer Technical Test

This test separate by 2 type test folder :
- Problem solving test
- ticket application (implementation test)


# Problem Solving Test
## How to run program

- Test code 1
run test1.go with following 5 number of argument input, example :

```bash
  go run test1.go 1 2 3 4 5
```

- Test code 2
run test2.go with input argument twice (length of array & data of array), example :

```bash
  go run test2.go
  6
  -4 3 -9 0 4 1
```

- Test code 3
run test1.go with following input argument time format, example :

```bash
  go run test3.go 07:05:45PM
```

# ticket application (implementation test)
## Installation

- Preparing library

```bash
  go mod tidy
```

- Database migration

```bash
  cd ticket_app
  go install -tags 'postgres,mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  migrate -database 'postgres://postgres:root@localhost:5432/ticket_app?sslmode=disable' -path migrations up
```

- Run server

```bash
  cd ticket_app
  go run server.go
```

## What this API could do ?

Complain ticket application

Auth API
- Login
- Register user
- Remove user

As a customer
- submit complaint ticket
- get ticket (filter by ticket_id)
- reply comment (filter by ticket_id)

As a supervisor
- fetch all ticket (optional filter by agent_id & ticket status)
- handover ticket to agent (filter by agent_id)
- get ticket (filter by ticket_id)
- Close ticket (filter by ticket_id)

As a agent
- fetch all ticket (filter by agent_id)
- get ticket (filter by ticket_id)
- reply comment (filter by ticket_id)
- Close ticket (filter by ticket_id)

Application flow
- create 1 supervisor and 1 agent
- user submit complaint ticket
- login as a supervisor
- supervisor fetch all ticket
- supervisor handover ticket to an agent
- login as and agent
- agent reply comment on a ticket
- close ticket