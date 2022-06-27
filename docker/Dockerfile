FROM golang:1.16.0-alpine3.13 as daemon
RUN go get github.com/githubnemo/CompileDaemon

FROM golang:1.18-alpine3.16 as dev
WORKDIR /app

COPY ./go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

COPY --from=daemon /go/bin/CompileDaemon /go/bin/ 

ENTRYPOINT [ "sh", "./docker/entrypoint.sh" ]
