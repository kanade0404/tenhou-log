FROM golang:1.18 AS builder
WORKDIR /app/
COPY .. .
WORKDIR /app/scraper
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/scraper/main .
EXPOSE 8088
ENV ENV=dev
CMD ["./main"]
