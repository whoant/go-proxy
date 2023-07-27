FROM golang:1.19-alpine3.16 as builder
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build -o main cmd/*.go

FROM alpine:3.16
WORKDIR /app

COPY --from=builder app/main .

EXPOSE 8080

CMD ["/app/main"]


