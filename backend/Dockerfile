FROM golang:alpine as builder

RUN apk add --no-cache git
RUN go get -u github.com/jinzhu/gorm 
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux

ENV SRC=/go/src/github.com/ComputePractice2018/photohosting/backend/
RUN mkdir -p '${SRC}'
ADD . ${SRC}


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/photohosting github.com/ComputePractice2018/photohosting/backend

FROM alpine:3.7

WORKDIR /opt
COPY --from=builder /go/bin/photohosting /opt/photohosting

EXPOSE 8080

ENV CONNECTION "photohosting:SuperSecretPassword@tcp(db:3306)/photohosting"
ENTRYPOINT /opt/photohosting -connection ${CONNECTION}