FROM golang:1.17-alpine3.15

WORKDIR /server/

COPY . .

RUN go build -o /bin/app .



FROM alpine:3.15

WORKDIR /server/

COPY --from=0 /bin/app ./app

EXPOSE 8001

CMD ["./app"]
