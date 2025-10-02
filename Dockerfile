#
# builder
#
FROM golang:1.24.0 AS builder

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /root/app

COPY . .

RUN go mod download
RUN go build -o /tmp/app

#
# runner
#
FROM gcr.io/distroless/static-debian12

COPY --from=builder /tmp/app /app

CMD ["/app"]
