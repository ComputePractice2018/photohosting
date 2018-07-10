FROM golang:alpine as builder

ENV SRC=/go/src/github.com/ComputePractice2018/photohosting
RUN mkdir -p '${SRC}'
ADD main.go ${SRC}
ADD utils/ ${SRC}

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/photohosting ${SRC}

FROM alpine:3.7

WORKDIR /opt
COPY --from=builder /go/bin/photohosting /opt/photohosting


ENTRYPOINT [ "/opt/photohosting" ]