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
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/MARIE .
```

If you want to use this backend inside a docker image, you have to set certificates into /etc/ssl/certs/ca-certificates.crt for API.ai SSL.