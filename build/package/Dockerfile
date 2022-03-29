# Build stage

FROM golang:1.16-alpine as build

# Add sqlite build dependencies
RUN apk add --update gcc musl-dev

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY scripts internal main.go ./

RUN sh scripts/build.sh

# Deploy stage

FROM alpine:3.15

WORKDIR /

COPY --from=build /app/bin/lnkshrt .

EXPOSE 8080

CMD /lnkshrt
