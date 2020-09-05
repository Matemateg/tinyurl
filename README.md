# TinyURL

This is an application that shortens URLs.  
Demo on heroku: [neznamovas-tinyurl.herokuapp.com](https://neznamovas-tinyurl.herokuapp.com/)

## Features
- [x] saving a short url
- [x] redirecting to original url
- [x] docker-compose
- [x] web UI
- [x] deployed to heroku, [click to open](https://neznamovas-tinyurl.herokuapp.com/)

## How to run
- `docker-compose up --build` - launching server on port 9000, database with schema
- `docker-compose up mysql` - launching only database with schema, but without application, 
application can be run `go run main.go`
- `docker-compose down` - to stop and remove all containers

## How to config
For config you can set environment variables:

- `PORT` - port HTTP server (default 9000)  
- `MYSQL_DSN` - dsn database (default dsn of database in docker)

The database schema is in the file `schema.sql`.

## How to use

#### Frontend
`open http://localhost:9000`

#### API
Create a tiny url: `curl http://localhost:9000/api/create -d '{"url": "https://avito.ru"}'`
