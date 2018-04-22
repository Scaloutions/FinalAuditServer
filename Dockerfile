FROM golang:1.9.2-alpine3.6 AS build

# Install tools required to build the project
RUN apk add --no-cache git

RUN mkdir -p /app

RUN go get "github.com/gin-gonic/gin"

RUN go get "gopkg.in/mgo.v2"

RUN go get -v "github.com/beevik/etree"

RUN go get "gopkg.in/mgo.v2/bson"

RUN go get "github.com/gorilla/mux"

RUN go get "github.com/golang/glog"

WORKDIR /app

ADD . /app

RUN go build -o server ./*.go 

CMD [ "/app/server", "-logtostderr=true"]

EXPOSE 8082
