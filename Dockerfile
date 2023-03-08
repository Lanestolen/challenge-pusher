FROM golang:1.19-alpine as builder
WORKDIR /app

COPY . .

RUN go build -o haaukins-exercises .

FROM alpine:3.16
COPY --from=builder /app/haaukins-exercises /

CMD ["./haaukins-exercises"]