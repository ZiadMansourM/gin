# 🫡 Welcome to examples-gin Using Go 🦦

# First Approach:
```Go
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	router.Run(":3000")
}
```

# Second Approach "Plug and Play":

```
ziadh@Ziads-MacBook-Air src % tree
.
├── Dockerfile.dev
├── Dockerfile.production
├── books
│   ├── controller.go
│   ├── models.go
│   ├── router.go
│   └── serializer.go
├── config
│   ├── db.go
│   └── settings.go
├── go.mod
├── go.sum
├── logs
│   └── log.log
├── main.go
├── middlewares
│   ├── auth.go
│   └── logger.go
└── users
    ├── controller.go
    ├── models.go
    ├── router.go
    └── serializer.go
```

### Vision: Reusability matters
This approach praise the reusability of apps, as in [Django docs](https://docs.djangoproject.com/en/4.0/intro/reusable-apps/#reusability-matters), It’s a lot of work to design, build, test and maintain a web application. Many projects share common problems. Wouldn’t it be great if we could save some of this repeated work?

### We will use Makefile to ease managing different environments
```makefile
THIS_FILE := $(lastword $(MAKEFILE_LIST))
dev_compose_file = docker-compose-dev.yml
staging_compose_file = docker-compose-staging.yml
production_compose_file = docker-compose-production.yml
.PHONY: help build up start down destroy stop restart logs logs-api ps login-timescale login-api db-shell
help:
	make -pRrq  -f $(THIS_FILE) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
production:
	docker-compose -f $(production_compose_file) up $(c)
staging:
	docker-compose -f $(staging_compose_file) build $(c)
build:
	docker-compose -f $(dev_compose_file) build $(c)
up:
	docker-compose -f $(dev_compose_file) up -d $(c)
reload:
	docker-compose -f $(dev_compose_file) down $(c)
	docker-compose -f $(dev_compose_file) up -d $(c)
ps:
	docker-compose -f $(dev_compose_file) ps
update:
	docker-compose -f $(dev_compose_file) up --build -d $(c)
start:
	docker-compose -f $(dev_compose_file) start $(c)
down:
	docker-compose -f $(dev_compose_file) down $(c)
destroy:
	docker-compose -f $(dev_compose_file) down -v $(c)
stop:
	docker-compose -f $(dev_compose_file) stop $(c)
restart:
	docker-compose -f $(dev_compose_file) stop $(c)
	docker-compose -f $(dev_compose_file) up -d $(c)
logs:
	docker-compose -f $(dev_compose_file) logs --tail=100 -f $(c)
logs-api:
	docker-compose -f $(dev_compose_file) logs --tail=100 -f api
login-timescale:
	docker-compose -f $(dev_compose_file) exec timescale /bin/bash
login-api:
	docker-compose -f $(dev_compose_file) exec api /bin/bash
db-shell:
	docker-compose -f $(dev_compose_file) exec timescale psql -Upostgres
```

### Dockerfile.dev
```Dockerfile.dev
FROM golang:1.18.2-alpine3.16 AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

ENTRYPOINT ["go", "run", "main.go"]
```

### docker-compose-dev.yml
```yml
version: '3'
services:
  db:
    image: postgres
    volumes:
      - ./data/db:/var/lib/postgresql/data
    ports:
      - "5432:5432"
    restart: always
    environment:
      - POSTGRES_USER=<secrets>
      - POSTGRES_PASSWORD=<secrets>
      - POSTGRES_DB=<secrets>
  backend:
    build:
      context: src
      dockerfile: Dockerfile.dev
    volumes:
      - ./src:/app
    ports:
      - "3000:3000"
    restart: always
    depends_on: 
      - db
```

### Dockerfile.production
```Dockerfile
FROM golang:1.18.2-alpine3.16 AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /server

FROM scratch
WORKDIR /app
COPY --from=builder /server .
EXPOSE 3000

ENTRYPOINT ["/app/server"]
```

### docker-compose-production.yml
```yml
version: '3'
services:
  db:
    image: postgres
    volumes:
      - ./data/db:/var/lib/postgresql/data
    restart: always
    environment:
      - POSTGRES_USER=<secrets>
      - POSTGRES_PASSWORD=<secrets>
      - POSTGRES_DB=<secrets>
  backend:
    build:
      context: src
      dockerfile: Dockerfile.production
    ports:
      - "3000:3000"
    restart: always
    depends_on: 
      - db
```

-----

Ziad Hassanin - SWE Doing SRE
-----------------------------
phone: (10)217-999-50 <br/>
ziadmansour.4.9.2000@gmail.com <br/>
San José State University, California <br/>
Cairo University Faculty of Engineering (CCEC) <br/>
[SREboy.com](https://www.sreboy.com/) | [twitter](https://twitter.com/ziad_m_404) | [linkedin](https://www.linkedin.com/in/ziad-mansour/) | [instagram](https://www.instagram.com/ziad_m_404/) <br/>