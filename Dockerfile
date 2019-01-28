FROM golang:1.11

ENV GOPATH /go

RUN go version

COPY . /go/src/KUMA-server/

RUN go get -u github.com/go-session/gin-session \
&& go get -u github.com/gin-gonic/gin \
&& go get -u github.com/jinzhu/gorm \
&& go get -u github.com/gin-gonic/gin \
&& go get -u github.com/lib/pq

WORKDIR /go/src/KUMA-server

CMD go run main.go

