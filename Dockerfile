# This file is a template, and might need editing before it works on your project.
FROM dockerhub.ir/golang:1.18 AS builder

WORKDIR /usr/src/app

COPY . .
RUN go get -v .
RUN go build -v -o app

FROM dockerhub.ir/buildpack-deps:buster as app

WORKDIR /usr/local/bin
RUN mkdir tmp
COPY --from=builder /usr/src/app/app .
#CMD ["./app", "$cmdrun"]
