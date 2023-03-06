FROM golang:1.20.1-alpine

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .
COPY .env .
RUN go mod download

RUN go build -o /efishery-auth cmd/main.go

EXPOSE 8888
CMD [ "/efishery-auth" ]
