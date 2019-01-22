FROM golang:1.11-stretch
LABEL author="Gustavo Henrique <gustavo@gustavohenrique.net>"

COPY . /go/src/github.com/gustavohenrique/grpc-golang-nodejs/catalog
WORKDIR /go/src/github.com/gustavohenrique/grpc-golang-nodejs/catalog

ENV PORT 11080

RUN go get google.golang.org/grpc \
 && go get github.com/golang/protobuf/proto \
 && go get github.com/gorilla/mux \
 && go get github.com/gorilla/handlers \
 && go get github.com/lib/pq \
 && go get github.com/jmoiron/sqlx

RUN sh -c "go build -i -v -o catalog main.go \
 && mv catalog /usr/local/bin"

CMD ["catalog"]
