#builder
FROM golang:1.17.7-alpine3.15 as builder

WORKDIR /app

# ARG db_host "default-value-arg"
# ENV DBHOST=${db_host}

COPY . .

RUN go mod download
RUN go build -o api app/main.go

#runner 
FROM alpine:3.15

COPY --from=builder /app/api /app/
# RUN apk add --no-cache bash

WORKDIR /app

EXPOSE 8080

CMD ["./api"]