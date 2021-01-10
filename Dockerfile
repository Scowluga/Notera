FROM golang:1.15-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app/server

RUN go build -o main .

EXPOSE $SERVER_PORT
CMD ["/app/server/main"]