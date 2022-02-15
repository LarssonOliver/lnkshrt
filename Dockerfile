# Build stage

FROM golang:1.16-alpine as build

ENV GOOS=linux
ENV GOARCH=arm64

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /lnkshrt

# Deploy stage

FROM alpine:3.15

RUN addgroup -S nonroot && adduser -S nonroot -G nonroot

WORKDIR /

COPY --from=build /lnkshrt .

USER nonroot:nonroot

EXPOSE 8080

CMD ["./lnkshrt"]
