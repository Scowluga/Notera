FROM golang:1.15-alpine  AS build

RUN mkdir /app
ADD . /app
WORKDIR /app/server

RUN go build -o main .

CMD ["/app/server/main"]