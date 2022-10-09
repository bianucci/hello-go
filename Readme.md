# Hello go!

## Overview

I started with a [basic go tutorial](https://go.dev/doc/tutorial/web-service-gin).

Extended it by an additional [auth layer](https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817).

I did no bother to set up a db so everything is in-memory and your users and albums are lost on re-start.

Token expiration is set to 20 seconds. See `token_lifespan_seconds`.

## Usage

See `client.rest` for sample requests.

You basically have to register a new user.

Log in and obtain a jwt.

All APIs under /api/albums are protected and can only be used by logged in users.

## Development

Set up

```bash
go get .
```

Single run

```bash
go run .
```

Watcher mode for convenience

```bash
npm install -g nodemon
```

```bash
nodemon --exec go run hello.go --signal SIGTERM
```
