# Atlas Corp. DNS Server

It is the Drone Navigation Service written in Go.

It only has one route (POST) and that is `/v1/dns` which is responsible for calculating the location from given coordinates and velocity.

Request body sample:

```json
{
  "x": "90.5",
  "y": "219",
  "z": "11",
  "vel": "15.6"
}
```

> Note: Be careful about the double quotes!

Response sample:

```json
{
  "loc": 3220.6
}
```

### Run tests

```shell script
go test ./...
```

### Run the application

First, you need to create `.env` file based on `.env.example`.

> Note: If you do not include DNS_PORT in the file, app will be served on 8080.

#### With docker and docker-compose

```shell script
docker-compose build
docker-compose up
``` 

#### Without docker

Be sure to install Go on your machine (1.12 preferred).

```shell script
go mod download
go run main.go
```