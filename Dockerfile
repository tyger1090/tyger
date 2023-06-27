# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

# Download Go modules
COPY /go_web_app .

RUN go build  .

# Run
CMD ["./news-demo-starter-files"]