FROM golang:alpine as builder

WORKDIR /app
COPY . .

RUN go build -o api .

FROM alpine:latest as runner
COPY --from=builder /app/api /api
WORKDIR /

ENTRYPOINT ["/api"]
