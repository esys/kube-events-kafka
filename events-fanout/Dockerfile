FROM golang:1.15.6-alpine3.12 as build
WORKDIR /build
COPY . .
RUN apk add alpine-sdk
RUN GOOS=linux GOARCH=amd64 go build -tags musl -o bin/events-fanout cmd/main.go

FROM golang:1.15.6-alpine3.12
EXPOSE 8080
WORKDIR /app
COPY --from=build /build/bin/events-fanout .
CMD ["./events-fanout"]