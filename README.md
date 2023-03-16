# scooby-api


## What is scooby-api?

scooby-api is a RESTful API for the scooby project. It is written in Golang using the Echo framework.

## How do I get set up?

### Dependencies

* [Golang](https://golang.org/)

### Installation

* `go get github.com/alexellis/scooby-api`
* `cd $GOPATH/src/github.com/alexellis/scooby-api`
* `go build`
* `./scooby-api`
* `curl localhost:8080`
* `curl localhost:8080/healthz`

### How to run tests

* `go test`
* `go test -cover`
* `go test -coverprofile=coverage.out`
* `go tool cover -html=coverage.out`

### How to run in Docker

* `docker build -t scooby-api .`
* `docker run -p 8080:8080 scooby-api`
* `curl localhost:8080`


###
```shell
docker build -t scoobychat.renatomoura.top:32000/scoobyapi:1 .
docker tag scoobychat.renatomoura.top:32000/scoobyapi:1 scoobychat.renatomoura.top:32000/scoobyapi:latest
docker push scoobychat.renatomoura.top:32000/scoobyapi:latest
kubectl apply -f ./k8s/
#kubectl rollout restart deployment/scoobyapi -n default
``