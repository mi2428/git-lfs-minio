FROM golang:1.11 as builder
LABEL maintainer="mi2428 <tmiya@protonmail.ch>"

WORKDIR /go/src/github.com/mi2428/git-lfs-minio
COPY . .

RUN go get -d -v github.com/minio/minio-go github.com/gorilla/mux
RUN  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/git-lfs-minio app.go

FROM alpine:latest

WORKDIR /opt/
COPY --from=builder /go/bin/git-lfs-minio .

EXPOSE 8080

CMD ["./git-lfs-minio"]
