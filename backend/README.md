# MARIE backend

> This backend is here to register new things and serve their data to other medias.

## Setup

You must have a mongodb process running

```shell
go get
go run main.go
```

You can also run this project with realize

```shell
go get github.com/tockins/realize
realize run
```

## Build

```shell
env GOOS=linux GOARCH=x86 go build -o ./bin/MARIE
```